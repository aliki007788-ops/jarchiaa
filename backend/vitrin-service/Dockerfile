FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o vitrin-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/vitrin-service .
EXPOSE 8090
CMD ["./vitrin-service"]


