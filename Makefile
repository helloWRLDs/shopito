#!make
include .env

path:
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin

grpc.init: path
	sudo apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
protoc:
	protoc --go-grpc_out=. ./services/**/protobuf/*.proto
	# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./services/products/protobuf/products.proto

docker.up:
	docker-compose up -d

docker.down:
	docker-compose down 

migrate-up:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

migrate-up-to:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up-to ${VOL}

migrate-down:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down

migrate-down-to:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down-to ${VOL}

goose.init:
	export PATH="$$PATH:$(GOPATH)/bin"
	go install github.com/pressly/goose/v3/cmd/goose@v3.19.2

users.run:
	go run ./services/users/cmd/app

notifier.run:
	go run ./services/notifier/cmd/app