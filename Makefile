.PHONY: help build run test lint clean install-tools

# Variables
BINARY_NAME=app
VERSION?=dev
BUILD_TIME=$(shell powershell -Command "Get-Date -Format 'yyyy-MM-dd HH:mm:ss'")
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X 'main.BuildTime=${BUILD_TIME}'"

help:
	@powershell -Command "Select-String -Path Makefile -Pattern '##' | ForEach-Object { $$_.Line }"

install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed successfully"

build:
	@echo "Building..."
	go build ${LDFLAGS} -o bin/${BINARY_NAME}.exe cmd/app/main.go
	@echo "Build complete: bin/${BINARY_NAME}.exe"

run:
	@echo "Running application..."
	go run cmd/app/main.go

test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	@echo "Tests complete"

test-coverage: test
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	@echo "Running linters..."
	golangci-lint run
	@echo "Linting complete"

lint-fix:
	@echo "Running linters with auto-fix..."
	golangci-lint run --fix

fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

vet:
	@echo "Running go vet..."
	go vet ./...

clean:
	@echo "Cleaning..."
	@if (Test-Path bin) { Remove-Item -Recurse -Force bin }
	@if (Test-Path coverage.out) { Remove-Item -Force coverage.out }
	@if (Test-Path coverage.html) { Remove-Item -Force coverage.html }
	@echo "Clean complete"

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod verify

tidy:
	@echo "Tidying dependencies..."
	go mod tidy

update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

docker-build:
	@echo "Building Docker image..."
	docker build -t ${BINARY_NAME}:${VERSION} .

all: clean deps lint test build
