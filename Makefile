
run:
	go run cmd/apiserver/main.go

build:
	go build -v ./cmd/apiserver

.DEFAULT_GOAL := build