
all: migrate-up build run

run:
	./apiserver

build:
	go build -v ./cmd/apiserver

migrate-down:
	migrate -database postgres://mark:123@localhost:5432/testdb?sslmode=disable -path migrations down

migrate-up:
	migrate -database postgres://mark:123@localhost:5432/testdb?sslmode=disable -path migrations up

.DEFAULT_GOAL := build