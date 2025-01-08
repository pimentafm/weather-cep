FROM golang:1.23.4-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weatherapi .

FROM scratch
WORKDIR /app
COPY --from=builder /app/weatherapi .
ENTRYPOINT [ "./weatherapi" ]