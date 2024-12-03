package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/smtp"
	"strings"
	"time"
	"unicode"

	"cloud.google.com/go/storage"                                      // GCP Storage 클라이언트 라이브러리
	_ "github.com/go-sql-driver/mysql"                                 // MySQL 드라이버
	historypb "github.com/mail-cote/go-server/history-service/history" // History 모듈의 gRPC 패키지
	mailpb "github.com/mail-cote/go-server/mail-service/mail"

	"google.golang.org/grpc"
)

const portNumber = "9000"

type mailServer struct {
	mailpb.MailServer
}

// User 구조체: DB에서 가져온 데이터를 담을 구조체
type Member struct {
	MemberId int64
	Email    string
	Level    string
}

// 로컬 MySQL 연결 정보
const (
	DBUser     = "root"        // MySQL 사용자 이름
	DBPassword = "gdsc1111"    // MySQL 비밀번호
	DBHost     = "34.22.95.16" // 로컬 MySQL 서버 (localhost)
	DBPort     = "3306"        // MySQL 포트
	DBName     = "mail_cote"   // MySQL 데이터베이스 이름
)

// GCP Storage 설정
const (
	BucketName = "mail-cote-bucket"
)

// SMTP 설정
const (
	SMTPServer   = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUsername = "mailcote1111@gmail.com"
	SMTPPassword = "zfmvzogpftiyqeeb"
)

// grpc: 버킷에서 랜덤 퀴즈값 가져오기
func (s *mailServer) FetchQuizFromBucket(ctx context.Context, req *mailpb.FetchQuizFromBucketRequest) (*mailpb.FetchQuizFromBucketResponse, error) {
	level := req.GetLevel()

	ctxGCP := context.Background()
	client, err := storage.NewClient(ctxGCP)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// level에서 마지막 숫자와 문자 분리
	var letters, digits string
	for i := len(level) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(level[i])) {
			digits = string(level[i]) + digits
		} else {
			letters = level[:i+1]
			break
		}
	}

	// "bronze"와 "5"로 분리된 값을 사용하여 폴더 경로 구성
	prefix := fmt.Sprintf("problems/%s/%s/", letters, digits)
	it := client.Bucket(BucketName).Objects(ctx, &storage.Query{Prefix: prefix})

	var files []string
	for {
		attrs, _ := it.Next()
		if attrs == nil {
			break
		}
		files = append(files, attrs.Name)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no files found in folder: %s", prefix)
	}

	rand.Seed(time.Now().UnixNano())
	randomFile := files[rand.Intn(len(files))]

	rc, err := client.Bucket(BucketName).Object(randomFile).NewReader(ctxGCP)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	content, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return &mailpb.FetchQuizFromBucketResponse{
		QuizContent: string(content),
		Message:     "Quiz successfully fetched!",
	}, nil
}

// grpc: 메일 보내기
func (s *mailServer) SendMail(ctx context.Context, req *mailpb.SendMailRequest) (*mailpb.SendMailResponse, error) {
	to := req.GetSendTo()
	from := req.GetSendFrom()
	quizContent := req.GetQuizContent()

	// JSON 데이터 파싱
	var data map[string]string
	err := json.Unmarshal([]byte(quizContent), &data)
	if err != nil {
		return nil, err
	}

	// HTML 템플릿 파일 경로
	templatePath := "mail_template.html"

	// HTML 템플릿 읽기 및 파싱
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	// 문제 설명 데이터 분리 및 HTTP 링크 감지
	descriptionLines := []struct {
		Content string
		IsImage bool
	}{}
	descriptionParts := strings.Split(data["description"], "\n") // 줄 단위로 분리
	for _, part := range descriptionParts {
		if strings.HasPrefix(part, "http://") || strings.HasPrefix(part, "https://") {
			descriptionLines = append(descriptionLines, struct {
				Content string
				IsImage bool
			}{Content: part, IsImage: true})
		} else {
			descriptionLines = append(descriptionLines, struct {
				Content string
				IsImage bool
			}{Content: part, IsImage: false})
		}
	}

	// 템플릿에 전달할 데이터 생성
	templateData := struct {
		QuizTitle        string
		Field            string
		TimeLimit        string
		MemoryLimit      string
		DescriptionLines []struct {
			Content string
			IsImage bool
		}
		InputDesc  string
		OutputDesc string
		InputEx    string
		OutputEx   string
		URL        string
	}{
		QuizTitle:        data["quiz_title"],
		Field:            data["field"],
		TimeLimit:        data["time_limit"],
		MemoryLimit:      data["memory_limit"],
		DescriptionLines: descriptionLines,
		InputDesc:        data["input_desc"],
		OutputDesc:       data["output_desc"],
		InputEx:          data["input_ex"],
		OutputEx:         data["output_ex"],
		URL:              data["url"],
	}

	// JSON 데이터를 템플릿에 바인딩
	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, templateData)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTML template: %v", err)
	}

	// SMTP 메일 전송 설정
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPServer)
	header := fmt.Sprintf("MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n")
	header += fmt.Sprintf("Subject: %s\r\n", from)
	header += fmt.Sprintf("To: %s\r\n", to)

	// HTML 템플릿을 이메일 본문으로 사용
	message := header + "\r\n" + bodyBuffer.String()

	err = smtp.SendMail(SMTPServer+":"+SMTPPort, auth, SMTPUsername, []string{to}, []byte(message))
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}

	return &mailpb.SendMailResponse{
		Message: "Email sent successfully!",
	}, nil
}

//**********************************************************************************************************************************

// connectToMySQL: SSH를 통해 MySQL 연결
func connectToMySQL() (*sql.DB, error) {
	// MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPassword, DBHost, DBName)

	// net.Dial을 사용하여 MySQL과 연결
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	return db, nil
}

// fetchUsersFromDB: GCP SQL에서 사용자 정보 가져오기
func fetchUsersFromDB(db *sql.DB) ([]Member, error) {
	rows, err := db.Query("SELECT member_id, email, level FROM Member")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		if err := rows.Scan(&member.MemberId, &member.Email, &member.Level); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

// ************************************************************************************************************************************
// 매일 아침 7시에 실행되는 작업
func dailyTask(s *mailServer, historyClient historypb.HistoryClient) {
	for {
		// // 현재 시간 확인
		// now := time.Now()

		// // 매일 아침 7시로 설정
		// nextRun := time.Date(now.Year(), now.Month(), now.Day(), 7, 0, 0, 0, now.Location())

		// // 현재 시간 이후 7시가 되기를 기다림
		// if now.After(nextRun) {
		// 	nextRun = nextRun.Add(24 * time.Hour) // 7시가 지났다면 다음날 7시로 설정
		// }

		// // 다음 실행 시간까지 대기
		// time.Sleep(nextRun.Sub(now))

		// 1분 대기
		time.Sleep(30 * time.Second)

		// 사용자별 작업 수행
		log.Println("Starting task for sending quizzes every minute...")

		// MySQL 연결
		db, err := connectToMySQL()
		if err != nil {
			log.Printf("Failed to connect to MySQL: %v", err)
			continue
		}
		defer db.Close()

		// 사용자 정보 가져오기
		users, err := fetchUsersFromDB(db)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			continue
		}

		// 각 사용자에 대해 JSON 파일을 랜덤으로 선택하고 메일 전송
		for _, user := range users {
			log.Printf("Processing user: %s (Level: %s)", user.Email, user.Level)

			// History 조회: 해당 사용자에게 보낸 퀴즈 ID 목록 가져오기
			historyReq := &historypb.GetAllHistoryRequest{UserId: user.MemberId}
			historyResp, err := historyClient.GetAllHistory(context.Background(), historyReq)
			if err != nil {
				log.Printf("Error fetching history for user %s: %v", user.Email, err)
				continue
			}

			// 보냈던 문제 ID 목록
			sentQuizIDs := map[int64]bool{}
			for _, quizID := range historyResp.QuizIds {
				sentQuizIDs[quizID] = true
			}

			// FetchQuizFromBucket 호출
			var selectedQuizContent string
			var selectedQuizId int64
			for attempt := 0; attempt < 5; attempt++ { // 최대 5번까지 중복되지 않는 문제를 시도
				quizResponse, err := s.FetchQuizFromBucket(context.Background(), &mailpb.FetchQuizFromBucketRequest{
					Level: user.Level,
				})
				if err != nil {
					log.Printf("Error fetching quiz for user %s: %v", user.Email, err)
					continue
				}

				// 퀴즈의 고유 ID 확인 (예: 파일 이름 또는 JSON 내 ID)
				quizID := quizResponse.GetQuizID()
				if !sentQuizIDs[quizID] {
					selectedQuizContent = quizResponse.GetQuizContent()
					selectedQuizId = quizResponse.GetQuizID()
					sentQuizIDs[quizID] = true
					break
				}
			}

			// selectedQuiz가 빈 문자열인 경우 (즉, 5번 모두 중복된 문제를 가져온 경우) 처리
			if selectedQuizContent == "" {
				log.Printf("No unique quiz could be fetched for user %s after 5 attempts", user.Email)
				continue
			}

			// SendMail 호출: 퀴즈 내용을 이메일로 전송
			sendMailResponse, err := s.SendMail(context.Background(), &mailpb.SendMailRequest{
				SendTo:      user.Email,
				SendFrom:    "mailcote1111@gmail.com", // 보낼 이메일 주소 설정
				QuizContent: selectedQuizContent,
			})
			if err != nil {
				log.Printf("Error sending mail to %s: %v", user.Email, err)
			} else {
				log.Printf("Mail sent to %s successfully! Response: %s", user.Email, sendMailResponse.Message)
			}

			// History에 전송 기록 저장
			historyAddReq := &historypb.SaveHistoryRequest{
				UserId: user.MemberId,
				QuizId: selectedQuizId, // 퀴즈의 고유 ID
			}
			_, err = historyClient.SaveHistory(context.Background(), historyAddReq)
			if err != nil {
				log.Printf("Error updating history for user %s: %v", user.Email, err)
			} else {
				log.Printf("History updated for user %s", user.Email)
			}
		}
	}
}

func main() {
	// gRPC 서버 시작
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mailServer := &mailServer{} // mailServer 객체 생성
	mailpb.RegisterMailServer(grpcServer, mailServer)

	// History 서비스의 gRPC 서버 주소
	const historyServiceAddress = "localhost:9002"

	// gRPC 연결 생성
	conn, err := grpc.Dial(historyServiceAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("🚨 Failed to connect to History service: %v", err)
	}
	defer conn.Close()

	// History 서비스 클라이언트 생성
	historyClient := historypb.NewHistoryClient(conn)

	go func() {
		log.Printf("start gRPC server on %s port", portNumber)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// 매일 아침 7시에 작업을 실행하는 goroutine 시작
	go dailyTask(mailServer, historyClient) // mailServer 객체를 dailyTask에 전달

	// 서버가 종료되지 않도록 대기
	select {}
}
