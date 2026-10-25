SHELL := /bin/bash
BUILD_DIR := _build

.PHONY: build/cmd
build/cmd:
	go build -o $(BUILD_DIR)/lorectl cmd/cmd.go

PHONY: build
build: build/cmd

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
