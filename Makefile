include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d postgres

env-down:
	@docker compose down postgres
env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных! (y/n) " ans;\
	if [ "$$ans" = "y" ]; then \
		docker compose down -v postgres port-forwarder && echo "Volume файлы очищены."; \
		rm -rf out/pgdata; \
		echo "Директория данных удалена."; \
	else \
		echo "Операция отменена."; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Ошибка: Не указано имя миграции. Используйте 'make migrate-create seq=имя_миграции'"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Ошибка: Не указано действие миграции. Используйте 'make migrate-action action=up|down'"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		"$(action)"

env-port-forward:
	@docker compose up -d port-forwarder
env-port-forward-close:
	@docker compose down port-forwarder

todoapp-stop:
	@PORT=$$(echo $(HTTP_SERVER_ADDR) | sed 's/^://'); \
	PIDS=$$(lsof -ti:$$PORT); \
	if [ -n "$$PIDS" ]; then \
		echo "Останавливаю процессы на порту $$PORT: $$PIDS"; \
		kill $$PIDS; \
		sleep 1; \
	else \
		echo "Порт $$PORT свободен."; \
	fi

todoapp-run: todoapp-stop
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	go mod tidy && \
	go run cmd/todoapp/main.go