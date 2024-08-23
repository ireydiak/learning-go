BINARY_NAME=go-todo

# Default
all: build

# Build the application
build:
	@echo "Building the binary..."
	GO111MODULE=on go build -o ./bin/$(BINARY_NAME) main.go

# Clean up build files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Install dependencies
deps:
	@echo "Building dependencies"
	GO111MODULE=on go mod tidy

# Install the binary
install: build
	@echo "Installing the binary to /usr/local/bin/$(BINARY_NAME)..."
	install -m 755 $(BINARY_NAME) /usr/local/bin

# Uninstall the binary
uninstall:
	@echo "Deleting /usr/local/bin/$(BINARY_NAME)..."
	rm -rf /usr/local/bin/$(BINARY_NAME)

# Format the code
format:
	@echo "Formatting the code..."
	go fmt ./...

.PHONY: all build clean deps install uninstall format
