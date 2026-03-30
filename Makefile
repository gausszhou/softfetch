.PHONY: build test clean install uninstall build-all lint fmt vet

BINARY_NAME=softfetch
DIST_DIR=dist

build:
	go build -o $(DIST_DIR)/$(BINARY_NAME) .

test:
	go test ./...

clean:
	rm -rf $(DIST_DIR)
	rm -f coverage.out

install:
	@echo "Installing $(BINARY_NAME)..."
	go install .
	@echo "$(BINARY_NAME) installed successfully to $$(go env GOPATH)/bin"

uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	rm -f $$(go env GOPATH)/bin/$(BINARY_NAME).exe
	@echo "$(BINARY_NAME) uninstalled successfully"

build-all: build-linux build-darwin build-windows

build-linux:
	mkdir -p $(DIST_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64 .

build-darwin:
	mkdir -p $(DIST_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 .

build-windows:
	mkdir -p $(DIST_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-arm64.exe .

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...