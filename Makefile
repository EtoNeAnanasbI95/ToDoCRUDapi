CONNECTION_STRING ?= $(PGCONNECT)
repo_dirs = tasks_repository_test users_repository_test
service_dirs = users_service_test

all: migrate run

migrate:
	# postgres://$(CRUD_USER):$(CRUD_PASSWORD)@localhost:5432/$(DBNAME)?sslmode=disable
	migrate -database "$(CONNECTION_STRING)" -path migrations up

rollback:
	migrate -database "$(CONNECTION_STRING)" -path migrations down

gen-swag:
	@echo "Generate swagger docs"
	@swag i -d ./cmd/ToDoCRUD/,./internal,./models
	@echo "Done"

run:
	@echo "Starting api"
	@go run ./cmd/ToDoCRUD/main.go --config="./config/local.yaml"

run-tests:
	@echo "Running tests..."
	$(foreach dir, $(repo_dirs), @TEST_DB_DSN=$(TEST_DB_DSN) go test "./tests/repositories_test/$(dir)" -v)
	$(foreach dir, $(service_dirs), @TEST_DB_DSN=$(TEST_DB_DSN) go test "./tests/services_test/$(dir)" -v)
