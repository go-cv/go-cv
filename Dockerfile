# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static" -X main.APP_VERSION=${APP_VERSION:-latest} -X main.COMMIT_ID=${COMMIT_ID:-undefined}' -o gocv .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/gocv .

# Copy config file
COPY config.yaml .

# Create directories for content and output
RUN mkdir -p /app/content /app/output /app/themes

# Copy themes
COPY --from=builder /app/themes ./themes

# Expose default port
EXPOSE 80

# Default command runs serve mode
ENTRYPOINT ["./gocv"]
CMD ["serve"]