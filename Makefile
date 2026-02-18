.PHONY: build testnet test clean help install docker-build

help:
	@echo "Pickle - Data Preservation Engine"
	@echo ""
	@echo "Usage:"
	@echo "  make build          Build pickle chain binary"
	@echo "  make install        Install pickle chain binary"
	@echo "  make testnet        Start single-validator testnet"
	@echo "  make test           Run tests"
	@echo "  make clean          Clean build artifacts"
	@echo "  make proto          Generate protocol buffers"
	@echo "  make docker-build   Build Docker image"

build:
	@echo "Building Pickle chain..."
	@go build -o ./bin/pickled ./cmd/pickled

install:
	@echo "Installing Pickle chain..."
	@go install ./cmd/pickled

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf ./bin
	@go clean

proto:
	@echo "Generating protocol buffers..."
	@buf generate

testnet:
	@echo "Starting single-validator testnet..."
	@./scripts/testnet.sh

docker-build:
	@echo "Building Docker image..."
	@docker build -t pickle:latest .

.DEFAULT_GOAL := help
