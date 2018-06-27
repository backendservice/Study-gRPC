# Study gRPC

Visual Studio Code

Go언어 설치
* https://golang.org/dl

Go Path 설정
* 환경변수 GOPATH = go 설치 폴더
* 환경변수 path에 작업 경로 추가 (%GOPATH%\bin)

gRPC
* https://grpc.io/
* 설치 - go get -u google.golang.org/grpc

ProtocolBuffers
* https://developers.google.com/protocol-buffers/
* 설치 - http://github.com/google/protobuf/releases/tag/v3.5.1

Protobuf Gen Go
* Grab the code from the repository and install the proto package.
  The simplest way is to run go get -u github.com/golang/protobuf/protoc-gen-go
* 설치 - go get -u github.com/golang/protobuf/protoc-gen-go

Dep
* https://github.com/golang/dep#installation
* 설치 - go get -u github.com/golang/dep/cmd/dep