.PHONY: dev build run test clean

# Variables
BINARY_NAME=tmp/main
CMD_PATH=./cmd/server

# Detect OS (Windows pakai .exe)
ifeq ($(OS),Windows_NT)
	BINARY_NAME := tmp/main.exe
endif

# Default target
dev:
	air

build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

run: build
	$(BINARY_NAME)

test:
	go test ./...

clean:
	rm -rf tmp/*

