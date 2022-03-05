build:
	go build -o bin/listener main.go

run:
	go run main.go

arm: 
	GOOS=linux GOARCH=arm64 go build -o bin/listener_arm64 

static:
	GOOS=linux GOARCH=amd64 go build -tags osusergo,netgo
	
all:	build arm

