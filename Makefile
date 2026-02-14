BINARY_NAME ?= api-server
DOCKER_IMAGE ?= go-nuts:latest
GO ?= go

.PHONY: help tidy fmt vet test build run run-metrics clean docker-build docker-run

help:
	@echo "Available targets:"
	@echo "  make tidy         - Run go mod tidy"
	@echo "  make fmt          - Format Go files"
	@echo "  make vet          - Run go vet"
	@echo "  make test         - Run tests"
	@echo "  make build        - Build api-server binary into ./bin"
	@echo "  make run          - Run api-server locally"
	@echo "  make run-metrics  - Run metrics-analyzer against README.md"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker image exposing :8080"

tidy:
	$(GO) mod tidy

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./...

build:
	mkdir -p bin
	CGO_ENABLED=0 $(GO) build -o bin/$(BINARY_NAME) ./cmd/api-server

run:
	$(GO) run ./cmd/api-server

run-metrics:
	$(GO) run ./cmd/metrics-analyzer ./README.md

clean:
	rm -rf bin

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run --rm -p 8080:8080 --name go-nuts $(DOCKER_IMAGE)
