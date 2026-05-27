# Builder
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build a static binary
# CGO_ENABLED=0 disables cgo, GOOS=linux ensures the binary is built for Linux, and -a -installsuffix cgo forces a static build
# disabling CGO ensures that it runs on any minimal Linux distribution
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-todo main.go


# Final minimal image
FROM alpine:latest

# Install root certs for secure HTTPS connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/go-todo .

# Expose the app port and the argus log port
EXPOSE 8080 9090

# Run the application
CMD ["./go-todo"]