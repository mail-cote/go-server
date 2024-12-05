package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"cloud.google.com/go/storage"                                      // GCP Storage 클라이언트 라이브러리
	_ "github.com/go-sql-driver/mysql"                                 // MySQL 드라이버
	historypb "github.com/mail-cote/go-server/history-service/history" // History 모듈의 gRPC 패키지
	mailpb "github.com/mail-cote/go-server/mail-service/mail"
	memberpb "github.com/mail-cote/go-server/member-service/member"

	"google.golang.org/grpc"
)

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
	SMTPPassword = "ldqnvppvbktsktee"
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

	// 파일 경로를 "/"로 분리하여 배열로 만든 후, 마지막 부분을 가져옵니다.
	parts := strings.Split(randomFile, "/")
	quizIdStr := parts[len(parts)-1] // 배열에서 마지막 부분을 가져옴 (예: "5087.json")

	quizIdStr = quizIdStr[:len(quizIdStr)-5] // ".json"을 제거
	// 숫자 문자열을 int64로 변환
	quizId, err := strconv.ParseInt(quizIdStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid quiz id in filename: %v", err)
	}

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
		QuizId:      quizId,
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

	// HTML 템플릿 읽기 및 파싱
	tmpl, err := template.ParseFiles("mail_template.html")
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
	header := fmt.Sprintf("MIME-version: 1.0\r\n")
	header += fmt.Sprintf("Content-Type: text/html; charset=\"UTF-8\";\r\n")
	header += "Subject: 오늘의 코딩 테스트 문제가 도착했어요!\r\n"
	header += fmt.Sprintf("To: %s\r\n", to)
	header += "\r\n" // 헤더와 본문을 구분하는 빈 줄 추가

	dialer := &net.Dialer{
		Timeout: 60 * time.Second, // 타임아웃을 설정합니다.
	}

	// HTML 템플릿을 이메일 본문으로 사용
	message := header + bodyBuffer.String()

	// SMTP 서버 연결
	conn, err := dialer.Dial("tcp", SMTPServer+":"+SMTPPort)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer conn.Close()

	// SMTP 클라이언트 생성
	client, err := smtp.NewClient(conn, SMTPServer)
	if err != nil {
		return nil, fmt.Errorf("failed to create SMTP client: %v", err)
	}
	// STARTTLS 연결 (암호화)
	if err := client.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
		return nil, fmt.Errorf("failed to start TLS: %v", err)
	}

	// SMTP 인증
	if err := client.Auth(auth); err != nil {
		return nil, fmt.Errorf("failed to authenticate: %v", err)
	}

	// 발신자 및 수신자 설정
	if err := client.Mail(from); err != nil {
		return nil, fmt.Errorf("failed to set MAIL FROM: %v", err)
	}
	if err := client.Rcpt(to); err != nil {
		return nil, fmt.Errorf("failed to set RCPT TO: %v", err)
	}

	// 메일 본문 작성
	writer, err := client.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get writer for email body: %v", err)
	}

	_, err = writer.Write([]byte(message))
	if err != nil {
		return nil, fmt.Errorf("failed to write email body: %v", err)
	}

	// 이메일 전송 완료
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close email writer: %v", err)
	}

	// SMTP 세션 종료
	if err := client.Quit(); err != nil {
		return nil, fmt.Errorf("failed to quit SMTP session: %v", err)
	}

	log.Printf("Mail sent to %s successfully!", to)

	// 성공적으로 메일을 보낸 후 응답을 반환
	return &mailpb.SendMailResponse{
		Message: "Email sent successfully!",
	}, nil
}

// 매일 아침 7시에 실행되는 작업
func dailyTask(s *mailServer, historyClient historypb.HistoryClient, memberClient memberpb.MemberServiceClient) {
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

		// 30초 대기
		time.Sleep(30 * time.Second)

		// 사용자별 작업 수행
		log.Println("Starting task for sending quizzes every minute...")

		memberReq := &memberpb.GetAllMemberRequest{}
		memberResp, err := memberClient.GetAllMember(context.Background(), memberReq)
		if err != nil {
			log.Fatalf("Failed to get members: %v", err)
		}

		var users []Member
		for _, grpcMember := range memberResp.Member {
			users = append(users, Member{
				MemberId: grpcMember.MemberId,
				Email:    grpcMember.Email,
				Level:    grpcMember.Level,
			})
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
				quizID := quizResponse.GetQuizId()
				log.Printf("Quiz number: %d", quizID)
				if !sentQuizIDs[quizID] {
					selectedQuizContent = quizResponse.GetQuizContent()
					selectedQuizId = quizResponse.GetQuizId()
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
				SendFrom:    SMTPUsername, // 보낼 이메일 주소 설정
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
				QuizId: selectedQuizId,
				Level:  user.Level,
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
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mailServer := &mailServer{} // mailServer 객체 생성
	mailpb.RegisterMailServer(grpcServer, mailServer)

	// gRPC 연결 생성(history)
	conn1, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("🚨 Failed to connect to History service: %v", err)
	}
	defer conn1.Close()

	// gRPC 연결 생성(member)
	conn2, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("🚨 Failed to connect to History service: %v", err)
	}
	defer conn2.Close()

	// History 서비스 클라이언트 생성
	historyClient := historypb.NewHistoryClient(conn1)
	memberServiceClient := memberpb.NewMemberServiceClient(conn2)

	go func() {
		log.Printf("start gRPC server on %s port", "9000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// 매일 아침 7시에 작업을 실행하는 goroutine 시작
	go dailyTask(mailServer, historyClient, memberServiceClient) // mailServer 객체를 dailyTask에 전달
	// 서버가 종료되지 않도록 대기
	select {}
}
