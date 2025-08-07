# Stage 1: Install dependencies
FROM golang:1.24.3-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Stage 2: Build the application
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY . .


RUN go build -ldflags="-w -s" -o main ./cmd/api/main.go

# Final stage: Run the application
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache shadow

# Create a non-root user and group
RUN groupadd -r appuser && useradd -r -g appuser appuser

# Copy the built application
COPY --from=builder /app/main .

# Change ownership of the application binary
RUN chown appuser:appuser /app/main

# Switch to the non-root user
USER appuser

CMD ["./main"]
