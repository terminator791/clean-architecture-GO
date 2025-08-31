.PHONY: build run test clean docker-build docker-up docker-down

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f server

# Build Docker image
docker-build:
	docker build -t clean-architecture-go .

# Start services with Docker Compose
docker-up:
	docker-compose up --build

# Stop Docker Compose services
docker-down:
	docker-compose down

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run migrations up
migrate-up:
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/products?sslmode=disable" up

# Run migrations down
migrate-down:
	migrate -path migrations -database "postgres://postgres:password@localhost:5432/products?sslmode=disable" down

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run