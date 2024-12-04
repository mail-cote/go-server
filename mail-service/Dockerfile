# 공식 Go 이미지를 Docker Hub에서 가져옴
FROM golang:1.23

# 컨테이너에서 /app 디렉터리를 컨테이너 내부의 작업 디렉터리로 설정
# 이후 모든 RUN, CMD, COPY는 이 디렉터리를 기준으로 실행
WORKDIR /app

# 프로토콜 버퍼 컴파일러 설치
RUN apt-get update && apt-get install -y protobuf-compiler

# 플러그인 설치 (각각 별도로 설치)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# $GOPATH/bin을 PATH에 추가
ENV PATH="/go/bin:${PATH}"

# 로컬(빌드 명령(.) 실행 위치) 기준
# go.mod와 go.sum 파일을 복사
# /app 디렉터리에 두 파일이 복사된다. 
COPY ./go.mod ./go.sum ./

# 디버깅: 컨테이너 내부 파일 구조 확인
RUN ls -R /app

# go.mod와 go.sum 파일을 기반으로 Go 모듈을 다운로드
RUN go mod download
RUN go mod tidy

# mail-service 디렉터리를 통째로 복사
COPY ./mail-service ./mail-service
COPY ./history-service ./history-service
COPY ./member-service ./member-service

# mail-service 디렉터리로 이동하여 gRPC 파일 생성
WORKDIR /app/mail-service    
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./mail/mail.proto

# 필요한 외부 모듈 설치
RUN go get github.com/go-sql-driver/mysql \
    github.com/joho/godotenv \
    google.golang.org/grpc
RUN go get github.com/mail-cote/go-server/history-service/history
RUN go get github.com/mail-cote/go-server/mail-service/mail 
RUN go get github.com/mail-cote/go-server/member-service/member 

RUN go get cloud.google.com/go/storage

# 컨테이너 내 /app/mail-service/server/main.go를 빌드하여 실행 가능한 바이너리 파일 server를 생성
RUN go build -o server_binary ./server/main.go

# 서버가 수신 대기할 포트 번호 설정
EXPOSE 9000

# 컨테이너가 실행될 때 /mail-service/server_binary 실행파일을 실행
CMD ["./server_binary"]