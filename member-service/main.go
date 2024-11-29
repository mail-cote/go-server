package memberservice

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/mail-cote/go-server/proto/member"
)

const (
	port = ":50052" // Member Service gRPC 포트
)

func main() {
	// TCP 리스너 설정
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC 서버 생성
	grpcServer := grpc.NewServer()

	// MemberService 서버 등록
	pb.RegisterMemberServiceServer(grpcServer, &MemberServiceServer{})

	log.Printf("Member Service is running on port %s", port)

	// 서버 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
