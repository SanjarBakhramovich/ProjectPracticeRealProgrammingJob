# Makefile для создания миграций

# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:password@localhost:5433/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Команда для запуска приложения
run:
	go run cmd/app/main.go
