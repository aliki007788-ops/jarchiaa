FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o analytics-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/analytics-service .
EXPOSE 8087
CMD ["./analytics-service"]


