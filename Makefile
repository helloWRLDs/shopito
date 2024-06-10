#!make
include .env

path:
	export GOPATH=$$HOME/go
	export PATH=$$PATH:$$GOPATH/bin

# Docker
docker.up:
	docker-compose up -d

docker.down:
	docker-compose down 

# Migrations
goose.init:
	export PATH="$$PATH:$(GOPATH)/bin"
	go install github.com/pressly/goose/v3/cmd/goose@v3.19.2

migrate-up:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

migrate-up-to:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up-to ${VOL}

migrate-down:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down

migrate-down-to:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down-to ${VOL}

# GRPC code gen
grpc.init: path
	sudo apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
protoc.example:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/products/protobuf/products.proto

protoc:
	find ./services -name "*.proto" -print0 | xargs -0 protoc --go_out=./pkg/protobuf --go-grpc_out=./pkg/protobuf

# sqlc code gen
sqlc:
	docker run --rm -v $$(pwd):/src -w /src sqlc/sqlc generate

# Swagger code gen
swag.api-gw:
	swag init -g ./services/api-gw/cmd/app/main.go -o ./services/api-gw/docs