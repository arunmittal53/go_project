# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
RUN go mod download

# 🔧 Cross-compile for Linux AMD64
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Minimal image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .

# Optional: install certs if needed for HTTPS
# RUN apk --no-cache add ca-certificates

EXPOSE 8080
ENTRYPOINT ["./main"]
