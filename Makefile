include .env
export

# Кроссплатформенное определение текущей директории
ifeq ($(OS),Windows_NT)
    export PROJECT_ROOT=$(CURDIR)
else
    export PROJECT_ROOT=$(shell pwd)
endif

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
ifeq ($(OS),Windows_NT)
	@powershell -Command " \
		[Console]::OutputEncoding = [System.Text.Encoding]::UTF8; \
		[Console]::InputEncoding = [System.Text.Encoding]::UTF8; \
		$$msg = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('0J7Rh9C40YHRgtC40YLRjCDQstGB0LUgdm9sdW1lINGE0LDQudC70Ysg0L7QutGA0YPQttC10L3QuNGPPyDQntC/0LDRgdC90L7RgdGC0Ywg0YPRgtC10YDQuCDQtNCw0L3QvdGL0YUuIFt5L05d')); \
		$$success = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('0KTQsNC50LvRiyDQvtC60YDRg9C20LXQvdC40Y8g0L7Rh9C40YnQtdC90Ys=')); \
		$$cancel = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('0J7Rh9C40YHRgtC60LAg0L7QutGA0YPQttC10L3QuNGPINC+0YLQvNC10L3QtdC90LA=')); \
		$$ans = Read-Host \"$$msg\"; \
		if ($$ans.Trim().ToLower() -eq 'y') { \
			docker compose down todoapp-postgres; \
			Remove-Item -Recurse -Force 'out/pgdata' -ErrorAction SilentlyContinue; \
			Write-Host $$success -ForegroundColor Green; \
		} else { \
			Write-Host $$cancel -ForegroundColor Gray; \
		}"
else
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi
endif

env-port-forward:
	@docker compose up -d  port-forwarder

env-port-close:
	@docker compose down port-forwarder

ifeq ($(OS),Windows_NT)
# === Версия для Windows ===
migrate-create:
	@powershell -Command " \
		[Console]::OutputEncoding = [System.Text.Encoding]::UTF8; \
		if ([string]::IsNullOrWhiteSpace('$(seq)')) { \
			$$err = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('0J7RgtGB0YPRgtGB0YLQstGD0LXRgiDQvdC10L7QsdGF0L7QtNC40LzRi9C5INC/0LDRgNCw0LzQtdGC0YAgc2VxLiDQn9GA0LjQvNC10YA6IG1ha2UgbWlncmF0ZS1jcmVhdGUgc2VxPWluaXQ=')); \
			Write-Host $$err -ForegroundColor Red; \
			exit 1; \
		}"
	@docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
else
# === Версия для Ubuntu / Linux ===
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1;\
	fi
	@docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
endif

# Таким образом можно динамически передавать какие-то различные значения внутрь Makefile таргета при их вызове
#test-target:
#	@echo "value: $(var)"
# команда для терминала: make test-target var=hello

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

ifeq ($(OS),Windows_NT)
# === Версия для Windows (PowerShell + Base64) ===
migrate-action:
	@powershell -Command " \
		[Console]::OutputEncoding = [System.Text.Encoding]::UTF8; \
		if ([string]::IsNullOrWhiteSpace('$(action)')) { \
			$$err = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('0J7RgtGB0YPRgtGB0YLQstGD0LXRgiDQvdC10L7QsdGF0L7QtNC40LzRi9C5INC/0LDRgNCw0LzQtdGC0YAgYWN0aW9uLiDQn9GA0LjQvNC10YA6IG1ha2UgbWlncmF0ZS1hY3Rpb24gYWN0aW9uPXVwL2Rvd24=')); \
			Write-Host $$err -ForegroundColor Red; \
			exit 1; \
		}"
	@docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
else
# === Версия для Ubuntu / Linux (Bash) ===
migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up/down"; \
		exit 1;\
	fi
	@docker compose run --rm todoapp-postgres-migrate \
		- path /migrations \
		- database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
endif