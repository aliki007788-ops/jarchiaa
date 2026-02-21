FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o geo-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/geo-service .
EXPOSE 8084
CMD ["./geo-service"]


