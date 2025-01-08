run:
	go run cmd/api/main.go

weatherapi-build:
	docker build -t "$(PROJECT_NAME)/weatherapi:v1.0.0" -f Dockerfile .

weatherapi-run:
	docker run -p 8080:8080 weatherapi

.PHONY: run
