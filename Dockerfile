# Stage 1: Base Environment (Shared tools and dependencies)
FROM golang:1.26-alpine AS base

# Install make for our Makefile commands
RUN apk add --no-cache make

# Optimization: Copy the pre-built linter directly from the official image
COPY --from=golangci/golangci-lint:latest /usr/bin/golangci-lint /usr/bin/golangci-lint

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .


# Stage 2: Testing Environment (Targeted by docker-compose for tests)
FROM base AS tester
# Run the strict linter first; if it passes, proceed with unit tests
CMD make lint && make test


# Stage 3: Builder (Compiles the application)
FROM base AS builder
# Build a static binary. CGO_ENABLED=0 disables cgo for maximum compatibility with minimal base images.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-todo main.go


# Stage 4: Production Release (Minimal image for deployment)
FROM alpine:latest AS release

# Install root certs for secure HTTPS connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/go-todo .

# Expose the app port and the argus telemetry port
EXPOSE 8080 9090

# Run the application
CMD ["./go-todo"]