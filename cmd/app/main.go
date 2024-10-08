package main

import (
	"REST/internal/database"
	"REST/internal/messagesService"
	"REST/internal/web/messages"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)
	
	// Инициализируем echo
	e := echo.New()
	
	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := messages.NewStrictHandler(handler, []messages.StrictMiddlewareFunc{}) // тут будет ошибка
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}