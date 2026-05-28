.PHONY: build run deps lint-simple test test-no-coverage \
 		lint clean docker-up docker-down coverage-summary \
		update-badge docker-test docker-lint

# Makefile for Go Todo App
# Provides commands for building, running, testing, and isolated CI tasks.

build:
	@echo "Building the Go Todo App..."
	go build -o bin/todo main.go

run: build
	@echo "Running the app..."
	./bin/todo

deps:
	@echo "Downloading and tidying dependencies..."
	go mod download
	go mod tidy

lint-simple:
	@echo "Running simple lint (vet, fmt)..."
	go vet ./...
	go fmt ./...

test:
	@echo "Running tests with coverage for all packages..."
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out
	@echo "Report coverage..."
	go tool cover -func=coverage.out

coverage-summary:
	@echo "Extracting total coverage..."
	@go tool cover -func=coverage.out | grep total | awk '{print $$3}'

# Dynamically updates the coverage badge in README.md without relying on 3rd party actions
update-badge:
	@echo "Updating coverage badge in README.md..."
	@COVERAGE=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}' | tr -d '%'); \
	sed -i -e "s/Coverage-[0-9.]*%25/Coverage-$${COVERAGE}%25/g" README.md

test-no-coverage:
	@echo "Running tests without coverage for faster execution..."
	go test ./... -v

lint:
	@echo "Running strict static analysis (golangci-lint)..."
	golangci-lint run ./...

clean:
	@echo "Cleaning up temporary files and coverage reports..."
	rm -rf bin/
	rm -rf logs/
	rm -rf *.log
	rm -f coverage.out

docker-up:
	@echo "Starting full stack with Docker..."
	docker compose up --build -d

docker-down:
	@echo "Stopping and removing Docker containers..."
	docker compose down -v

# Runs the isolated test service alongside the database
docker-test:
	@echo "Running tests inside an isolated Docker container..."
	docker compose up --build --abort-on-container-exit --exit-code-from test test
	@echo "Extracting coverage report from the test container..."
	docker compose cp test:/app/coverage.out ./coverage.out

docker-lint:
	@echo "Running golangci-lint in an isolated Docker container..."
	docker compose run --rm --build test make lint