BINARY_NAME := bin/isthere
LOCAL_INSTALL_DIR := $(HOME)/.local/bin
INSTALL_DIR := /usr/bin

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/isthere ./cmd/isthere

install-local: build
	@mkdir -p $(LOCAL_INSTALL_DIR)
	@cp $(BINARY_NAME) $(LOCAL_INSTALL_DIR)
	@echo "Installed $(BINARY_NAME) to $(LOCAL_INSTALL_DIR)"

clean:
	@rm -f $(BINARY_NAME)
	@echo "Cleaned up"

test:
	@echo "Running tests"
	@go test
