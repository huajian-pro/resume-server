all:
	make default
	make linux

default:
	go buind -o resume main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o resume-lin main.go

docker:
	make linux
	docker buildx build -t xiwuou/resume:0.1 --platform=linux/arm64,linux/amd64 . --push

