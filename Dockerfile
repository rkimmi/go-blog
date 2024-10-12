# Use the official Golang image for building
ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire application code to the container
COPY . .

# Build the application and output the binary as '/blog'
RUN go build -v -o /blog .

# Use a minimal image for the final stage (Debian Bookworm)
FROM debian:bookworm

# Copy the binary from the builder stage
COPY --from=builder /blog /usr/local/bin/

# Expose the port your app listens on, if needed
EXPOSE 8080

# Set the command to run the binary
CMD ["/usr/local/bin/blog"]