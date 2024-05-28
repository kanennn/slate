BINARY_NAME=slate
BUILD_PATH=bin
MAIN=cmd/main.go

build: build-all

run:
	@go run $(MAIN)

build-local:
	@go build -o $(BUILD_PATH)/$(BINARY_NAME)-local $(MAIN)

build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-windows-amd64 build-windows-arm64 

build-minimal: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-windows-amd64

build-darwin-amd64:
	@GOARCH=amd64 GOOS=darwin go build -o $(BUILD_PATH)/$(BINARY_NAME)-darwin-amd64 $(MAIN) 

build-darwin-arm64:
	@GOARCH=arm64 GOOS=darwin go build -o $(BUILD_PATH)/$(BINARY_NAME)-darwin-arm64 $(MAIN) 
	
build-linux-amd64:
	@GOARCH=amd64 GOOS=linux go build -o $(BUILD_PATH)/$(BINARY_NAME)-linux-amd64 $(MAIN) 
	
build-linux-arm64:
	@GOARCH=arm64 GOOS=linux go build -o $(BUILD_PATH)/$(BINARY_NAME)-linux-arm64 $(MAIN)

build-windows-amd64:
	@GOARCH=amd64 GOOS=windows go build -o $(BUILD_PATH)/$(BINARY_NAME)-windows-amd64 $(MAIN) 
	
build-windows-arm64:
	@GOARCH=arm64 GOOS=windows go build -o $(BUILD_PATH)/$(BINARY_NAME)-windows-arm64 $(MAIN) 
