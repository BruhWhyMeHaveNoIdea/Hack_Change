# Multi-stage build for assignment-service
FROM golang:1.24.10-alpine AS builder

# Install git for go modules that require it
RUN apk add --no-cache git

WORKDIR /src

# Copy go.mod and go.sum to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the sources
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w" -o /app ./cmd/main

# Runtime image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /
COPY --from=builder /app /app

EXPOSE 8080
ENTRYPOINT ["/app"]
