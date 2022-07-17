
all: migrate-up build run

run:
	./apiserver

build:
	go build -v ./cmd/apiserver

migrate-down:
	migrate -database postgres://mark:123@localhost:5432/testdb?sslmode=disable -path migrations down

migrate-up:
	migrate -database postgres://mark:123@localhost:5432/testdb?sslmode=disable -path migrations up

post-req:
	http http://localhost:8081/blob type="article" title="idk how to do that task help" link_self="none" link_related="none" author_type="alien" author_id=1

.DEFAULT_GOAL := build