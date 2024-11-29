package problemservice

import (
	"context"
	"fmt"
	"io/ioutil"

	pb "github.com/mail-cote/go-server/proto" // Protobuf 파일에서 생성된 패키지 경로

	"google.golang.org/protobuf/encoding/protojson" // JSON 변환 라이브러리
)

// ProblemServiceServer는 gRPC 서버의 구현체
type ProblemServiceServer struct {
	pb.UnimplementedProblemServiceServer
}

// GetProblemByNumber은 문제 번호로 문제를 가져오는 API
func (s *ProblemServiceServer) GetProblemByNumber(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemResponse, error) {
	// 문제 파일 경로 수정하기 !!!!!!!!!!!!!!!
	filePath := fmt.Sprintf("problems/%s.json", req.QuizNum)

	// 파일 읽기
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read problem file: %w", err)
	}

	// JSON 파일을 Protobuf 형식으로 변환
	var problem pb.Problem
	if err := protojson.Unmarshal(data, &problem); err != nil {
		return nil, fmt.Errorf("failed to unmarshal problem JSON: %w", err)
	}

	return &pb.GetProblemResponse{Problem: &problem}, nil
}
