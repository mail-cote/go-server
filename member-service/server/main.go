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
	query := "INSERT INTO Member (email, level, password) VALUES (?, ?, ?)"
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
	query := "UPDATE Member SET level = ?, password = ? WHERE member_id = ?"
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
	query := "DELETE FROM Member WHERE member_id = ?"
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
// 1. CreateMember 테스트
/*
func testInsertData(db *sql.DB) {
	query := "INSERT INTO Member (email, password, level) VALUES (?, ?, ?)"
	_, err := db.Exec(query, "testuser@example.com", "password", "silver2")
	if err != nil {
		log.Fatalf("Failed to insert test data: %v", err)
	}
	log.Println("✅ Test data inserted successfully!")
}

// 2. UpdateMember 테스트
func testUpdateMember(s *MemberServiceServer) {
	// 테스트 데이터 준비
	insertQuery := "INSERT INTO Member (email, level, password) VALUES (?, ?, ?)"
	result, err := s.db.Exec(insertQuery, "updatetest@example.com", "bronze3", "oldpassword")
	if err != nil {
		log.Fatalf("🚨 Failed to insert test data: %v", err)
	}

	// 삽입된 데이터의 ID 확인
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("🚨 Failed to retrieve last insert ID: %v", err)
	}

	// UpdateMember 요청 생성
	req := &pb.UpdateMemberRequest{
		MemberId: int32(lastInsertID), // int를 string으로 변환
		Level:    "gold2",
		Password: "newpassword",
	}

	// UpdateMember 호출
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := s.UpdateMember(ctx, req)
	if err != nil {
		log.Fatalf("🚨 UpdateMember failed: %v", err)
	}

	log.Printf("✅ UpdateMember response: %s", resp.Message)
}

func testDeleteMember(s *MemberServiceServer) {
	// 테스트 데이터 준비
	insertQuery := "INSERT INTO Member (email, level, password) VALUES (?, ?, ?)"
	result, err := s.db.Exec(insertQuery, "deletetest@example.com", "silver1", "password")
	if err != nil {
		log.Fatalf("🚨 Failed to insert test data: %v", err)
	}

	// 삽입된 데이터의 ID 확인
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("🚨 Failed to retrieve last insert ID: %v", err)
	}

	// DeleteMember 요청 생성
	req := &pb.DeleteMemberRequest{
		MemberId: int32(lastInsertID), // int를 string으로 변환
	}

	// DeleteMember 호출
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := s.DeleteMember(ctx, req)
	if err != nil {
		log.Fatalf("🚨 DeleteMember failed: %v", err)
	}

	log.Printf("✅ DeleteMember response: %s", resp.Message)
}
*/

func main() {
	// TCP 리스너 설정
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC 서버 생성
	grpcServer := grpc.NewServer()

	// MemberService 서버 초기화
	server := NewMemberServiceServer()
	defer server.db.Close() // 서버 종료 시 DB 연결 닫기

	// CreateMember 테스트
	// testInsertData(server.db)

	// UpdateMember 테스트
	// testUpdateMember(server)

	// DeleteMember 테스트
	// testDeleteMember(server)

	pb.RegisterMemberServiceServer(grpcServer, server)

	log.Printf("Member Service is running on port %s", port)

	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
