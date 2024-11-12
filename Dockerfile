# 멀티스테이지 빌드: Dockerfile을 빌드 단계와 런타임 단계로 구분한 것
# 1단계: Go 어플리케이션 생성
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
# main.go가 cmd에 위치함
RUN go build -o myapp ./cmd/main.go  

# 2단계: 앱을 실행하기 위한 최소 이미지의 준비
FROM gcr.io/distroless/base
COPY --from=builder /app/myapp /myapp
CMD ["/myapp"]
