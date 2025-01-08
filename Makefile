include .env

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weatherapi cmd/api/main.go

run:
	go run cmd/api/main.go

weatherapi-build:
	docker build -t "$(PROJECT_NAME)/weatherapi:v1.0.0" -f Dockerfile .

weatherapi-run:
	docker run --name ${PROJECT_NAME}-weatherapi -p 8080:8080 -e  WEATHERAPI_API_KEY=${WEATHERAPI_API_KEY} "$(PROJECT_NAME)/weatherapi:v1.0.0"

.PHONY: run
