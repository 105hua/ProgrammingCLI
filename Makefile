.PHONY: build clean run test test-verbose test-coverage test-coverage-report docker-test docker-test-coverage docker-test-race docker-clean help

# Binary name
BINARY_NAME=ProgrammingCLI
BUILD_DIR=build

# Build the application
build:
	@echo Building $(BINARY_NAME)...
	@go build -o $(BUILD_DIR)/$(BINARY_NAME).exe main.go
	@echo Build complete: $(BUILD_DIR)/$(BINARY_NAME).exe

# Build for multiple platforms
build-all:
	@echo Building for multiple platforms...
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go
	@echo Build complete for all platforms in $(BUILD_DIR)/

# Run the application
run: build
	@echo Running $(BINARY_NAME)...
	@./$(BUILD_DIR)/$(BINARY_NAME).exe

# Clean build artifacts
clean:
	@echo Cleaning build directory...
	@rm -rf $(BUILD_DIR)
	@echo Clean complete

# Run tests
test:
	@echo Running tests...
	@go test ./...

# Run tests with verbose output
test-verbose:
	@echo Running tests with verbose output...
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo Running tests with coverage...
	@go test -cover ./...

# Run tests with detailed coverage report
test-coverage-report:
	@echo Generating coverage report...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo Coverage report generated: coverage.html

# Docker-based testing targets
docker-test:
	@echo Running tests in Docker container...
	@docker compose run --rm test

docker-test-coverage:
	@echo Running tests with coverage in Docker container...
	@docker compose run --rm test-coverage

docker-test-race:
	@echo Running tests with race detection in Docker container...
	@docker compose run --rm test-race

docker-clean:
	@echo Cleaning Docker images and volumes...
	@docker compose down -v
	@docker rmi programmingcli-test 2>/dev/null || true
	@echo Docker cleanup complete

# Display help
help:
	@echo Available targets:
	@echo   build                - Build the application for current platform
	@echo   build-all            - Build the application for all platforms
	@echo   run                  - Build and run the application
	@echo   test                 - Run all tests
	@echo   test-verbose         - Run tests with verbose output
	@echo   test-coverage        - Run tests with coverage statistics
	@echo   test-coverage-report - Generate HTML coverage report
	@echo   docker-test          - Run tests in Docker container
	@echo   docker-test-coverage - Run tests with coverage in Docker
	@echo   docker-test-race     - Run tests with race detection in Docker
	@echo   docker-clean         - Clean Docker images and volumes
	@echo   clean                - Remove build artifacts
	@echo   help                 - Display this help message
