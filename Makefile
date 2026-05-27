.PHONY: build run test clean docker-up docker-down

build:
	@echo "building the Go Todo App..."
	go build -o bin/todo main.go

run: build
	@echo "Running the app..."
	./bin/todo

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning up..."
	rm -rf bin/
	rm -rf logs/

docker-up:
	@echo "Starting full stack with Docker..."
	docker-compose up --build -d

docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down