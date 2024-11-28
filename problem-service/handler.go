package main

import (
	context "context"

	pb "go-server/proto" // 프로토버프 패키지 경로 다시 확인하기
)

type ProblemServiceServer struct {
	pb.UnimplementedProblemServiceServer
}

func (s *ProblemServiceServer) GetProblemByNumber(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemResponse, error) {
	// GCS에서 문제 파일 읽어오기
	/*
		problem := &pb.Problem{
			QuizNum:     req.QuizNum,
			Title:       "",
			Description: "",
			Field:       "",
		}
	*/
	return &pb.GetProblemResponse{Problem: problem}, nil
}
