all:
	make default
	make linux

run:
	go run main.go

run-prod:
	go run main.go -s=true >> logs/resume_log.out

default:
	go buind -o resume main.go
	chmod +x resume

linux:
	GOOS=linux GOARCH=amd64 go build -o resume main.go

docker-loc:
	make linux
	docker build -t resume:0.1 .

docker-hub:
	make linux
	docker buildx build -t xiwuou/resume:0.1 --platform=linux/arm64,linux/amd64 . --push