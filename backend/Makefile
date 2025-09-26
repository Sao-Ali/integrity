.PHONY: tidy lint build run test

tidy:
	go mod tidy

lint:
	golangci-lint run ./...

build:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

test:
	go test ./... -race

