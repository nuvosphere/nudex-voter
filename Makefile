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

.PHONY: test
test: ## go test
	go test  -short -v ./...

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

.PHONY: fmt
fmt: ## go fmt
	go fmt ./...
	gofumpt -l -w .
	gci write  -s standard -s default .

.PHONY: fix
fix: fmt ## auto fix code
	wsl --fix ./...
	tagalign -fix -sort ./...
	godot -w ./

.PHONY: dep
tool: ## Install dep tool
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install github.com/daixiang0/gci@latest
	go install github.com/bombsimon/wsl/v4/cmd/wsl@latest
	go install github.com/tetafro/godot/cmd/godot@latest
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/4meepo/tagalign/cmd/tagalign@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

.PHONY: abi
abi: ## generate abi go file
	abigen --abi internal/layer2/abis/erc20.json --pkg contracts --type ERC20 --out internal/layer2/contracts/erc20.go
	abigen --abi internal/layer2/abis/ierc1271.json --pkg contracts --type IERC1271 --out internal/layer2/contracts/ierc1271.go
	abigen --abi internal/layer2/abis/erc721.json --pkg contracts --type ERC721 --out internal/layer2/contracts/erc721.go
	abigen --abi internal/layer2/abis/erc1155.json --pkg contracts --type ERC1155 --out internal/layer2/contracts/erc1155.go
	abigen --abi internal/layer2/abis/multicall3.json --pkg contracts --type Multicall3 --out internal/layer2/contracts/multicall3.go
	abigen --abi internal/layer2/abis/AccountManager.json --pkg contracts --type AccountManagerContract --out internal/layer2/contracts/account_manager.go
	abigen --abi internal/layer2/abis/TaskManager.json --pkg contracts --type TaskManagerContract --out internal/layer2/contracts/task_manager.go
	abigen --abi internal/layer2/abis/AssetManager.json --pkg contracts --type AssetManagerContract --out internal/layer2/contracts/asset_manager.go
	abigen --abi internal/layer2/abis/ParticipantManager.json --pkg contracts --type ParticipantManagerContract --out internal/layer2/contracts/participant_manager.go
	abigen --abi internal/layer2/abis/VotingManager.json --pkg contracts --type VotingManagerContract --out internal/layer2/contracts/voting_manager.go
	abigen --abi internal/layer2/abis/DepositManager.json --pkg contracts --type DepositManagerContract --out internal/layer2/contracts/deposit_manager.go
	abigen --abi internal/layer2/abis/TaskPayload.json --pkg contracts --type TaskPayloadContract --out internal/layer2/contracts/task_payload.go

.PHONY: ci
ci: abi fix fmt lint build test 
	echo $? && echo "success!"

.PHONY: lint
lint: ## Runs the linter
	golangci-lint run ./...

.PHONY: help
help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	
.PHONY: all build clean test deps run docker-build