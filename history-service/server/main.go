package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버 예시

	historypb "github.com/mail-cote/go-server/history-service/history"
	"google.golang.org/grpc"
)

const port = ":9001"

type historyServer struct {
	historypb.HistoryServer
	DB *sql.DB
}

func getDBSource() string {
	dbUser := "root"         // 사용자명
	dbPassword := "gdsc1111" // 비밀번호
	dbHost := "34.22.95.16"  // 데이터베이스 호스트
	dbPort := "3306"         // 데이터베이스 포트
	dbName := "mail_cote"    // 데이터베이스 이름

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func NewHistoryServer() *historyServer {
	// MySQL 연결
	db, err := sql.Open("mysql", getDBSource()) // dbSource는 MySQL 정보

	if err != nil {
		log.Fatalf("🚨 Failed to connect to database: %v", err)
	}

	// DB 연결 확인
	if err := db.Ping(); err != nil {
		log.Fatalf("🚨 Database is unreachable: %v", err)
	}

	log.Println("Database connection successful!")

	return &historyServer{DB: db}
}

// 1. history 조회 -> userid가 같은 내용만
func (s *historyServer) getAllHistory(ctx context.Context, req *historypb.GetAllHistoryRequest) (*historypb.GetAllHistoryResponse, error) {
	if req == nil {
		return nil, errors.New("🚨 request object is nil")
	}
	if s.DB == nil {
		return nil, fmt.Errorf("🚨 database connection is not initialized")
	}
	userId := req.GetUserId()
	if userId == 0 {
		return nil, errors.New("🚨 Data is required")
	}

	query := "SELECT quiz_id FROM History WHERE member_id = ?"
	rows, err := s.DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("🚨 Failed to select histories: %v", err)
	}
	defer rows.Close()

	// quizIds를 저장할 슬라이스
	var quizIds []int64
	// 결과 읽기
	for rows.Next() {
		var quizId int64
		if err := rows.Scan(&quizId); err != nil {
			return nil, fmt.Errorf("🚨 failed to scan row: %w", err)
		}
		quizIds = append(quizIds, quizId)
	}

	// 에러가 발생했는지 체크
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("🚨 rows iteration error: %w", err)
	}

	// 응답 생성
	response := &historypb.GetAllHistoryResponse{
		QuizIds: quizIds,
	}
	return response, nil
}

// 2. history 저장
func (s *historyServer) saveHistory(ctx context.Context, req *historypb.SaveHistoryRequest) (*historypb.SaveHistoryResponse, error) {
	userId := req.GetUserId()
	quizId := req.GetQuizId()
	level := req.GetLevel()

	if userId == 0 || quizId == 0 || level == "" {
		return nil, errors.New("🚨 Data is required")
	}

	// 데이터베이스 쿼리
	query := "INSERT INTO History (user_id, quiz_id, level) VALUES (?, ?, ?)"
	_, err := s.DB.Exec(query, userId, quizId, level)
	if err != nil {
		return nil, fmt.Errorf("🚨 Failed to create history: %v", err)
	}

	return &historypb.SaveHistoryResponse{}, nil
}

func main() {
	// TCP 리스너 설정
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC 서버 생성
	grpcServer := grpc.NewServer()

	// MemberService 서버 초기화
	server := NewHistoryServer()
	defer server.DB.Close() // 서버 종료 시 DB 연결 닫기

	log.Printf("History Service is running on port %s", port)

	historypb.RegisterHistoryServer(grpcServer, server)

	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
