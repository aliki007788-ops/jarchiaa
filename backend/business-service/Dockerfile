FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o business-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/business-service .
EXPOSE 8082
CMD ["./business-service"]


