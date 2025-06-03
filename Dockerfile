# Start from official Go image
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api/main.go

# Runtime image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/deploy/migrations ./migrations

EXPOSE 8080

CMD ["./server"]
