# Variables
BINARY_NAME = number-guessing-game
BUILD_DIR = bin
SOURCE_DIR = ./cmd

# Build the binary
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Run game
run: build
	./${BUILD_DIR}/${BINARY_NAME}

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)