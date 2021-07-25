.DEFAULT_GOAL := default

default: build test

build:
	go build .

test:
	go test -v ./...


.PHONY: build test
