FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o admin-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/admin-service .
EXPOSE 8089
CMD ["./admin-service"]


