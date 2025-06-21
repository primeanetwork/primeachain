# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: geth all test clean devtools docker

GOBIN = ./build/bin
GO ?= latest
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_COMMIT_DATE=$(shell git log -n1 --pretty='format:%cd' --date=format:'%Y%m%d')

geth:
	mkdir -p $(GOBIN)
	go build -o $(GOBIN)/geth ./cmd/geth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."

all: geth

test:
	go test ./...

clean:
	env GO111MODULE=on go clean -cache
	rm -rf build/bin/*

devtools:
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/fjl/gencodec@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

docker:
	docker build --pull -t bnb-chain/bsc:latest -f Dockerfile .
