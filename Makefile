NAME = bsc-evm-compatible-bridge-api
ENTRYPOINT = ./cmd/bsc-evm-compatible-bridge-api-server
PACKAGES=$(shell go list ./... |grep -v client)
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TAGS = netgo
BUILD_FLAGS = #-tags "${BUILD_TAGS}"

all: format build

########################################
### CI

ci: build

########################################
### Generate

generate:
	swagger generate server -f ./swagger.yml -A bsc-evm-compatible-bridge-api --default-scheme=http

########################################
### Build

build:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/$(NAME).exe $(ENTRYPOINT)
else
	go build $(BUILD_FLAGS) -o build/$(NAME) $(ENTRYPOINT)
endif

build-linux:
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install:
	go install $(BUILD_FLAGS) $(ENTRYPOINT)

run:
	go run $(ENTRYPOINT) --scheme=http --port=8080

docker-build:
	docker build . -t bsc-evm-compatible-bridge-api
	
docker-run:
	docker run -it -p 8080:8080 --name bsc-evm-compatible-bridge-api --rm bsc-evm-compatible-bridge-api

docker-run-detach:
	docker run -dit -p 8080:8080 --name bsc-evm-compatible-bridge-api --rm bsc-evm-compatible-bridge-api

docker-stop:
	docker stop bsc-evm-compatible-bridge-api

########################################
### Format
format:
	@echo "--> Formatting"
	$(shell cd ../../../ && goimports -w -local github.com/binance-chain/$(NAME) $(PACKAGES))
	$(shell cd ../../../ && gofmt -w $(PACKAGES))

########################################
### Lint
lint:
	@echo "--> Lint"
	golint $(PACKAGES)

########################################
### Testing

test:
	make set_with_deadlock
	make test_unit
	make cleanup_after_test_with_deadlock

test_unit:
	@echo "--> go test -race"
	@go test -race $(PACKAGES) -v

########################################
### Pre Commit
pre_commit: build test format

########################################

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build install test test_unit build-linux
