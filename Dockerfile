# Use the official Golang image as the base
FROM golang:1.26.3-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source
COPY . .

# Build the Go binary
RUN go build -o /tmp/naas ./cmd/app/

# Final lightweight stage
FROM alpine:3.21 AS final

# Run from the app directory so reasons.json is found relative to CWD
WORKDIR /bin/app

# Copy the compiled binary from the builder stage
COPY --from=builder /tmp/naas /bin/app/naas

# Reasons data
COPY --from=builder /build/reasons.json /bin/app/reasons.json

# Expose the application's port
EXPOSE 8000

# Run the application
CMD ["./naas"]