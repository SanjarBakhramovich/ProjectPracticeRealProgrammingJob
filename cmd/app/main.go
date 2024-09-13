package main

import (
	"REST/internal/database"
	"REST/internal/handlers"
	"REST/internal/messagesService"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)




func main() {
	// Инициализация базы данных
	database.InitDB()
	// Автоматическая миграция модели Message
	database.DB.AutoMigrate(&messagesService.Message{})


	// Инициализация сервиса
	repo := messagesService.NewMessageRepository(database.DB)	// Инициализация обработчиков
	service := messagesService.NewService(repo)
	// Инициализация обработчиков
	handler := handlers.NewHandler(service)
	// ROUTERS
	// Создание маршрутизатора
	router := mux.NewRouter()
	// Установка маршрутов и связанных обработчиков
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	// 


	// Запуск сервера
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
