CONNECTION_STRING ?= $(PGCONNECT)
dirs = tasks_repository_test users_repository_test

all: migrate run

check-consts:
	ifeq ($(CONNECTION_STRING),)
	  $(error PGCONNECT не задан. Установите переменную окружения PGUSER или передайте её через make.)
	endif

migrate: check-consts
	# postgres://$(CRUD_USER):$(CRUD_PASSWORD)@localhost:5432/$(DBNAME)?sslmode=disable
	migrate -database "$(CONNECTION_STRING)" -path migrations up

rollback: check-consts
	migrate -database "$(CONNECTION_STRING)" -path migrations down

run:
	# TODO: написать запуск докера
	echo "приложение запускается"

run-tests:
	@echo "Running tests..."
	$(foreach dir, $(dirs), @TEST_DB_DSN=$(TEST_DB_DSN) go test "./tests/$(dir)" -v)

clean:
	@unset PG_USER PG_PASSWORD
	@echo "Переменные подключения очищены."
