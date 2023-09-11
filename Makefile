.PHONY: create-migration compose/up compose/down migrate/up migrate/status

create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "Please provide a migration name example: make create-migration NAME=create_users_table"; \
	else \
		docker run --rm -v $(PWD)/db/migrations:/migrations migrate/migrate:v4.14.1 create -ext sql -dir /migrations -seq $(NAME); \
	fi

migrate/up:
	docker compose up migrate

compose/up:
	docker compose up -d

compose/down:
	docker compose down -v # -v removes volumes

gql_gen:
	gqlgen generate