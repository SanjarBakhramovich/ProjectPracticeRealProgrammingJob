package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message = "global variable"

type MessageRequest struct {
	Message string `json:"message"`
}

func main() {
	// Инициализация базы данных
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	// Создание маршрутизатора
	router := mux.NewRouter()
	// Установка маршрутов и связанных обработчиков
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/update-message", UpdateMessageHandler).Methods("POST")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")

	// Запуск сервера
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
