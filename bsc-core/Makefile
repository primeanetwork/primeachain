# PrimeaChain Makefile (fixed for cmd/primea structure)
# Run from inside bsc-core

.PHONY: all geth test clean

GOBIN = ./build/bin

geth:
	mkdir -p $(GOBIN)
	go build -o $(GOBIN)/geth ./cmd/primea
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch Primea node."

all: geth

test:
	go test ./...

clean:
	rm -rf $(GOBIN)/*
