include .env

db_generate:
	@sqlc generate

db_migration:
	@migrate create -ext=sql -dir=internal/database/migrations -seq init

db_up:
	@migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

db_down:
	@migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

.PHONY: db_migration db_up db_down

# templ generate
# ./tailwindcss -i static/css/base.css -o static/css/tailwind.css --minify
