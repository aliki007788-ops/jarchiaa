FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o eitaa-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/eitaa-service .
EXPOSE 8088
CMD ["./eitaa-service"]


