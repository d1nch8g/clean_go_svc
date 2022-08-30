CURR_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
.PHONY: gen

# RUN DB FOR LOCAL DEVELOPMENT
db:
	docker run -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db -p 5432:5432 -d postgres

# RESTART SERVICE
run:
	docker compose down
	docker compose up --build app

# GENERATE PROTO AND SQLC CODE
gen:
	docker run --rm -v ${CURR_DIR}:/src -w /src rvolosatovs/protoc --proto_path=/src --go_out=. --go-grpc_out=. users.proto
	docker run --rm -v ${CURR_DIR}:/src -w /src kjconroy/sqlc generate -f sqlc.yml

# PUSH TO DOCKER HUB
push:
	docker build -t dangdancheg/clean_svc:latest .
	docker push dangdancheg/clean_svc:latest
	docker image remove dangdancheg/clean_svc:latest
