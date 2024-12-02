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

	"cloud.google.com/go/storage"      // GCP Storage 클라이언트 라이브러리
	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버
	mailpb "github.com/mail-cote/go-server/mail-service/mail"
	"google.golang.org/grpc"
)

const portNumber = "9000"

type mailServer struct {
	mailpb.MailServer
}

// User 구조체: DB에서 가져온 데이터를 담을 구조체
type Member struct {
	Email string
	Level string
}

// 로컬 MySQL 연결 정보
const (
	DBUser     = "tgwing"    // MySQL 사용자 이름
	DBPassword = "tgwing"    // MySQL 비밀번호
	DBHost     = "127.0.0.1" // 로컬 MySQL 서버 (localhost)
	DBPort     = "3306"      // MySQL 포트
	DBName     = "mail_cote" // MySQL 데이터베이스 이름
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPassword, "127.0.0.1", DBName)

	// net.Dial을 사용하여 MySQL과 연결
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	return db, nil
}

// fetchUsersFromDB: GCP SQL에서 사용자 정보 가져오기
func fetchUsersFromDB(db *sql.DB) ([]Member, error) {
	rows, err := db.Query("SELECT email, level FROM Member")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		if err := rows.Scan(&member.Email, &member.Level); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

//************************************************************************************************************************************

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mailpb.RegisterMailServer(grpcServer, &mailServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
