# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o customer-service ./cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/customer-service .

EXPOSE 8081

CMD ["./customer-service"]
