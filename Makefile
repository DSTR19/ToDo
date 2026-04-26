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
		docker compose down -v postgres && echo "Volume файлы очищены."; \
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