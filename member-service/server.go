package memberservice

import (
	"context"
	"errors"
	"log"

	pb "github.com/mail-cote/go-server/proto/member"
)

type MemberServiceServer struct {
	pb.UnimplementedMemberServiceServer
}

// CreateMember: 새 구독자 생성
func (s *MemberServiceServer) CreateMember(ctx context.Context, req *pb.CreateMemberRequest) (*pb.CreateMemberResponse, error) {
	// 요청 데이터 처리 로직 (예: DB에 저장)
	log.Printf("Creating member with email: %s", req.Email)

	// 예제 응답
	return &pb.CreateMemberResponse{
		MemberId: "generated-id",
		Message:  "Member created successfully",
	}, nil
}

// UpdateMember: 구독자 정보 업데이트
func (s *MemberServiceServer) UpdateMember(ctx context.Context, req *pb.UpdateMemberRequest) (*pb.UpdateMemberResponse, error) {
	// 요청 데이터 검증 및 처리
	if req.MemberId == "" {
		return nil, errors.New("Member ID is required")
	}

	log.Printf("Updating member %s with field: %s", req.MemberId, req.Field)

	// 예제 응답
	return &pb.UpdateMemberResponse{
		Message: "Member updated successfully",
	}, nil
}

// DeleteMember: 구독자 삭제
func (s *MemberServiceServer) DeleteMember(ctx context.Context, req *pb.DeleteMemberRequest) (*pb.DeleteMemberResponse, error) {
	// 요청 데이터 처리
	log.Printf("Deleting member %s", req.MemberId)

	// 예제 응답
	return &pb.DeleteMemberResponse{
		Message: "Member deleted successfully",
	}, nil
}
