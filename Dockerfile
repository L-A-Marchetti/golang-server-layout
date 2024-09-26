# Step 1: Build the application
FROM golang:1.22.5-bullseye AS builder

# Install SQLite3 development files
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

WORKDIR /app
COPY . .

# Download all dependencies
RUN find . -name go.mod -execdir go mod download \;

# Build the application with CGO enabled for SQLite support
RUN CGO_ENABLED=1 GOOS=linux go build -o golang-server-layout ./cmd/golang-server-layout

# Step 2: Create the final image
FROM debian:bullseye-slim

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3 \
    libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/golang-server-layout .

# Copy any other necessary files (e.g., templates, static files)
COPY --from=builder /app/web /root/web
COPY --from=builder /app/static /root/static

# Ensure the binary is executable
RUN chmod +x /root/golang-server-layout

EXPOSE 8080

CMD ["./golang-server-layout"]