# Makefile for nudex_voter

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=nudex-voter

# Build binary
all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

deps:
	$(GOGET) -u ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd && ./$(BINARY_NAME)

docker-build-all:
	docker buildx build --platform linux/amd64,linux/arm64 -t nuvosphere/nudex-voter:latest --push .

docker-build:
	docker buildx build --platform linux/amd64 -t nuvosphere/nudex-voter:latest --load .

docker-build-x:
	docker buildx build --platform linux/arm64 -t nuvosphere/nudex-voter:latest --load .

.PHONY: all build clean test deps run docker-build