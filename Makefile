.PHONY: build clean run help

# Binary name
BINARY_NAME=ProgrammingCLI
BUILD_DIR=build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME).exe main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME).exe"

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go
	@echo "Build complete for all platforms in $(BUILD_DIR)/"

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME).exe

# Clean build artifacts
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Display help
help:
	@echo "Available targets:"
	@echo "  build      - Build the application for current platform"
	@echo "  build-all  - Build the application for all platforms"
	@echo "  run        - Build and run the application"
	@echo "  clean      - Remove build artifacts"
	@echo "  help       - Display this help message"
