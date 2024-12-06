# Output directory
OUTPUT_DIR=releases

# Go source file
SOURCE_FILE=main.go

.PHONY: build-all windows macos macos-arm64 linux clean

build-all: windows macos macos-arm64 linux

windows: $(OUTPUT_DIR)
	@echo "Compiling for Windows (x86_64)..."
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/gh-open-windows.exe $(SOURCE_FILE)

macos: $(OUTPUT_DIR)
	@echo "Compiling for macOS (x86_64)..."
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/gh-open-macos $(SOURCE_FILE)

macos-arm64: $(OUTPUT_DIR)
	@echo "Compiling for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/gh-open-macos-arm64 $(SOURCE_FILE)

linux: $(OUTPUT_DIR)
	@echo "Compiling for Linux (x86_64)..."
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/gh-open-linux $(SOURCE_FILE)

clean:
	@echo "Cleaning up binaries..."
	rm -rf $(OUTPUT_DIR)
