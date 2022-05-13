install:
	go mod tidy

gen_grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/proto/api.proto

build:
	go build -o bin/server ./src/main.go

static:
	CGO_ENABLED=0 GOOS=windows go build -a -ldflags '-extldflags "-static"' -o server.exe ./src/main.go

start:
	go run ./src/main.go

all:
	go mod tidy && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/proto/api.proto && go build -o bin/server ./src/main.go && CGO_ENABLED=0 GOOS=windows go build -a -ldflags '-extldflags "-static"' -o server.exe ./src/main.go && go run ./src/main.go