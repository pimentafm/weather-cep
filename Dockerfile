FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weatherapi cmd/api/main.go

FROM alpine:3.21.2

WORKDIR /app
COPY .env /app

COPY --from=builder /app/weatherapi .

EXPOSE 8080

ENTRYPOINT [ "./weatherapi" ]