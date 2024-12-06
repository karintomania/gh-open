# Output directory
OUTPUT_DIR=releases

# Go source file
SOURCE_FILE=main.go

VERSION=$(shell grep "const version" main.go | sed -E 's/const version = "(.*)"/\1/')

.PHONY: build-all windows macos macos-arm64 linux clean print-version

print-version:
	@echo $(VERSION)

build-all: windows macos macos-arm64 linux

windows: $(OUTPUT_DIR)
	@echo "Compiling for Windows (x86_64)..."
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/gh-open-$(VERSION)-windows.exe $(SOURCE_FILE)

macos: $(OUTPUT_DIR)
	@echo "Compiling for macOS (x86_64)..."
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/gh-open-$(VERSION)-macos $(SOURCE_FILE)

macos-arm64: $(OUTPUT_DIR)
	@echo "Compiling for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/gh-open-$(VERSION)-macos-arm64 $(SOURCE_FILE)

linux: $(OUTPUT_DIR)
	@echo "Compiling for Linux (x86_64)..."
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/gh-open-$(VERSION)-linux $(SOURCE_FILE)

clean:
	@echo "Cleaning up binaries..."
	rm -rf $(OUTPUT_DIR)
