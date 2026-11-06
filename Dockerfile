# ==========================================
# Stage 1: Build stage
# ==========================================
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Install SSL certificates and timezone data
RUN apk add --no-cache ca-certificates tzdata

# Cache Go modules layer
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Compile a static, stripped binary (pure Go SQLite driver does not need CGO)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/server .

# ==========================================
# Stage 2: Minimal production runtime
# ==========================================
FROM alpine:3.21

WORKDIR /app

# Copy root CA certificates & timezone data from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Create non-root user and persistent storage directory for SQLite
RUN addgroup -S appgroup && adduser -S appuser -G appgroup && \
    mkdir -p /data && chown -R appuser:appgroup /data /app

# Copy the compiled binary from the builder
COPY --from=builder --chown=appuser:appgroup /app/server /app/server

# Default production environment variables
ENV PORT=8080 \
    DB_PATH=/data/books.db

# Expose API port
EXPOSE 8080

# Switch to non-root user for security
USER appuser

# Run binary
ENTRYPOINT ["/app/server"]

