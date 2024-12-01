package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/mail-cote/go-server/member-service/member"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const port = ":50052" // gRPC 포트

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getDBSource() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

// MemberServiceServer 구조체 정의
type MemberServiceServer struct {
	pb.UnimplementedMemberServiceServer
	db *sql.DB
}

// NewMemberServiceServer: 서버 초기화
func NewMemberServiceServer() *MemberServiceServer {
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

	return &MemberServiceServer{db: db}
}

// 기능1. CreateMember: 새 회원 생성
func (s *MemberServiceServer) CreateMember(ctx context.Context, req *pb.CreateMemberRequest) (*pb.CreateMemberResponse, error) {
	member := req.GetMember()
	if member == nil {
		return nil, errors.New("🚨 Member data is required")
	}

	// 입력값 검증
	if member.Email == "" || member.Password == "" {
		return nil, errors.New("🚨 Email and Password are required")
	}

	// 데이터베이스 쿼리
	query := "INSERT INTO member (email, level, password) VALUES (?, ?, ?)"
	_, err := s.db.Exec(query, member.Email, member.Level, member.Password)
	if err != nil {
		return nil, fmt.Errorf("🚨 Failed to create member: %v", err)
	}

	return &pb.CreateMemberResponse{
		Message: "Member created successfully",
	}, nil
}

// 기능2. UpdateMember: 회원 정보 업데이트
func (s *MemberServiceServer) UpdateMember(ctx context.Context, req *pb.UpdateMemberRequest) (*pb.UpdateMemberResponse, error) {
	query := "UPDATE member SET level = ?, password = ? WHERE member_id = ?"
	result, err := s.db.Exec(query, req.Level, req.Password, req.MemberId)
	if err != nil {
		return nil, fmt.Errorf("🚨 Failed to update member: %v", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, errors.New("🚨 No member found with the given ID")
	}

	return &pb.UpdateMemberResponse{
		Message: "Member updated successfully",
	}, nil
}

// 기능3. DeleteMember: 회원 삭제
func (s *MemberServiceServer) DeleteMember(ctx context.Context, req *pb.DeleteMemberRequest) (*pb.DeleteMemberResponse, error) {
	query := "DELETE FROM member WHERE member_id = ?"
	result, err := s.db.Exec(query, req.MemberId)
	if err != nil {
		return nil, fmt.Errorf("🚨 Failed to delete member: %v", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, errors.New("🚨 No member found with the given ID")
	}

	return &pb.DeleteMemberResponse{
		Message: "Member deleted successfully",
	}, nil
}

// ******************* 클라이언트 테스트 *****************************
// Member 테이블에 데이터를 삽입하는 테스트
func testInsertData(db *sql.DB) {
	query := "INSERT INTO Member (email, password, level) VALUES (?, ?, ?)"
	_, err := db.Exec(query, "testuser@example.com", "securepassword", "gold")
	if err != nil {
		log.Fatalf("Failed to insert test data: %v", err)
	}
	log.Println("Test data inserted successfully!")
}

func main() {
	// TCP 리스너 설정
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC 서버 생성
	grpcServer := grpc.NewServer()

	// MemberService 서버 등록
	server := NewMemberServiceServer()
	defer server.db.Close() // 서버 종료 시 DB 연결 닫기

	// 테스트 데이터 삽입
	testInsertData(server.db)

	pb.RegisterMemberServiceServer(grpcServer, server)

	log.Printf("Member Service is running on port %s", port)

	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
