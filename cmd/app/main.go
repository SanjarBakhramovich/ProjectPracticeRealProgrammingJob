package main

import (
	"REST/internal/database"
	"REST/internal/handlers"
	"REST/internal/messagesService"
	"log"
	"net/http"

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
	
	// Передаем и регистрируем хендлер в echo
	e.GET("api/messages", func(c echo.Context) error {
		response, err := handler.GetMessagesHandler(c.Request().Context(), messagesService.Message{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})
	
	e.POST("api/messages", func(c echo.Context) error {
		var req messagesService.Message
		if err := c.Bind(&req); err != nil {
			return err
		}

		response, err := handler.PostMessageHandler(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, response)
	})
	e.PATCH("api/patch/:id", handler.PatchMessageHandler) // Fixed PATCH route
	
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
