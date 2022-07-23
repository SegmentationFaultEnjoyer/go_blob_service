
all: migrate-up build run

gen: 
	./generate.sh -i $(pwd)/docs/web_deploy/ -p $(pwd)/resources

run:
	./apiserver

build:
	go build -v ./cmd/apiserver

migrate-down:
	migrate -database postgres://mark:123@localhost:5432/testdb?sslmode=disable -path migrations down

migrate-up:
	migrate -database postgres://mark:123@localhost:5432/testdb?sslmode=disable -path migrations up

req-get:
	http http://localhost:8081/integrations/testService/$(url)

req-del:
	http DELETE http://localhost:8081/integrations/testService/$(url)

req-post:
	http POST http://localhost:8081/integrations/testService/$(url) $(args)

post-req:
	http http://localhost:8081/integrations/testService/blob type="article" title="idk how to do that task help" link_self="none" link_next="none" link_prev="idk" link_first="idk" link_last="idk" author_type="mamka" author_id=2

.DEFAULT_GOAL := build