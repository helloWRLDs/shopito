#!make
include .env

set-path:
	export GOPATH=$$HOME/projects
	export PATH=$$PATH:/usr/bin:/bin

run:
	gnome-terminal --title="api" -- go run ./cmd/api

migrate-new:
	goose -dir migrations create ${NAME} ${TYPE}

goose-init: set-path
	go install github.com/pressly/goose/v3/cmd/goose@v3.19.2

migrate-up:
	goose -dir migrations ${DB} "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

migrate-up-to:
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up-to ${VOL}

migrate-down:
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down

migrate-down-to:
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down-to ${VOL}

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down 

start: docker-up run