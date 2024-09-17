# Makefile для создания миграций

# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:password@localhost:5433/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate-up:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Команда для запуска приложения
run:
	go run cmd/app/main.go

# Команда для генерации кода API из OpenAPI спецификации и запуска приложения
gen:
	oapi-codegen -config openapi/.openapi -include-tags messages -package messagesService openapi/openapi.yaml > ./internal/messagesService/api.gen.go


# Команда для облегчения git commit и push
git:
	@read -p "Введите сообщение коммита: " msg; \
	git add .; \
	git commit -m "$$msg"; \
	git push -u origin main
