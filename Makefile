install:
	go mod tidy
	echo "Installed all dependencies"

gen_grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/proto/api.proto
	echo "Generated Proto"

build:
	go build -o bin/server ./src/main.go
	echo "Built binary file"

static:
	CGO_ENABLED=0 GOOS=windows go build -a -ldflags '-extldflags "-static"' -o server.exe ./src/main.go
	echo "Built .exe"

start:
	echo "Starting server"
	go run ./src/main.go

git:
	git add .
	echo 'Input the commit message'
	read message
	git commit -m "$message"
	git push origin main

all:
	go mod tidy && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    src/proto/api.proto && go build -o bin/server ./src/main.go && CGO_ENABLED=0 GOOS=windows go build -a -ldflags '-extldflags "-static"' -o server.exe ./src/main.go && go run ./src/main.go