# Stage 1: Build the Go binary
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Copy dependency manifests and download layers
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked binary optimized for size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp ./app/

# Stage 2: Create the minimal runtime container
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Run the binary
CMD ["./myapp", "hello"]
