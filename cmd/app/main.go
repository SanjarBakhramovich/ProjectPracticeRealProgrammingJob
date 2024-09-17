package main

import (
	"REST/internal/database"
	"REST/internal/handlers"
	"REST/internal/messagesService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    // Инициализация базы данных
    database.InitDB()
    // Автоматическая миграция модели Message
    database.DB.AutoMigrate(&messagesService.Message{})

    // Инициализация сервиса
    repo := messagesService.NewMessageRepository(database.DB)
    service := messagesService.NewService(repo)
    // Инициализация обработчиков
    handler := handlers.NewHandler(service)

    // Инициализация сервера
    e := echo.New()

    // Используем Logger и Recover
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Передаем и регистрируем хендлеры в echo
    messagesHandler := messagesService.NewStrictHandler(handler, nil)
    messagesService.RegisterHandlers(e, messagesHandler)

    // Запуск сервера
    if err := e.Start(":8080"); err != nil {
        log.Fatalf("failed to start with err: %v", err)
    }
}
