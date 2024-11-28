# Base 이미지
FROM golang:1.20

# 작업 디렉토리 설정
WORKDIR /app

# Go 모듈 초기화 및 의존성 복사
COPY go.mod go.sum ./
RUN go mod download

# 서비스 코드 복사
COPY . .

# 프로토파일에서 생성된 코드 포함하여 빌드
RUN go build -o problem-service ./main.go

# gRPC 서버 실행
CMD ["./problem-service"]
