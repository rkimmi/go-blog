FROM golang:1.20-alpine AS builder

# Install build dependencies, including git (if needed for Go modules)
RUN apk add --no-cache git

# Set up the build environment
WORKDIR /usr/src/app

# Copy go.mod and go.sum to install dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code and build it
COPY . .
RUN go build -o /app/blog .

# Final stage: Use Alpine for the runtime environment
FROM alpine:latest

# Install CA certificates to verify HTTPS connections
RUN apk add --no-cache ca-certificates

# Set working directory and copy the Go binary from the builder stage
WORKDIR /root/
COPY --from=builder /app/blog .

# Expose the application port and run the binary
EXPOSE 8080
CMD ["./blog"]