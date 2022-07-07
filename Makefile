PWD := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# DEPENDENCY FIXES
dep:
	go get -u all
	go mod tidy

# RESTART SERVICE
run:
	docker compose down
	docker compose up --build app

# GENERATE PROTO AND SQLC CODE
generate:
	docker run --rm -v ${PWD}:/src -w /src rvolosatovs/protoc --proto_path=/src --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out . users.proto
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate -f sqlc.yml

# PUSH TO DOCKER HUB
push:
	docker build -t dangdancheg/clean_svc:latest .
	docker push dangdancheg/clean_svc:latest
	docker image remove dangdancheg/clean_svc:latest
