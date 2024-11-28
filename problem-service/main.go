package main

import (
	"log"
	"net"

	"github.com/mail-cote/go-server/problem-service/server" // server.go 파일 경로

	pb "github.com/mail-cote/go-server/proto/problem" // Protobuf 파일에서 생성된 패키지 경로

	"google.golang.org/grpc"
)

const (
	port = ":50051" // gRPC 서버 포트
)

func main() {
	// TCP 리스너 설정
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC 서버 생성
	grpcServer := grpc.NewServer()

	// ProblemService 서버 등록
	pb.RegisterProblemServiceServer(grpcServer, &server.ProblemServiceServer{})

	log.Printf("Problem Service is running on port %s", port)

	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
