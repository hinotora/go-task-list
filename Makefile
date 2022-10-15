.PHONY: build
build:
	go build -v ./cmd/go-task-list/main.go

.DEFAULT_GOAL := build