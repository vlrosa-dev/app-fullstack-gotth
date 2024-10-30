include .env

all: tailwind templ build run

run:
	@go run cmd/main.go

build:
	@go build -o bin/main cmd/main.go

templ:
	# Generate templ
	@templ generate

tailwind:
	# Generate css minify
	@./tailwindcss -i static/css/base.css -o static/css/tailwind.css -c ./tailwind.config.js --minify

db_generate:
	# Sqlc generate
	@sqlc generate

db_migration:
	# Generate migrations db
	@migrate create -ext=sql -dir=internal/database/migrations -seq init

db_up:
	# Start migrations up
	@migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

db_down:
	# Down migrations
	@migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

db_force:
	# Force migrations version 1
	@migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" force 1

.PHONY: db_generate db_migration db_up db_down db_force

# templ generate
# ./tailwindcss -i static/css/base.css -o static/css/tailwind.css --minify
