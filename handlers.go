package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Обработчик для создания сообщения
func CreateMessage(w http.ResponseWriter, r *http.Request) {
    var msg Message
    if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Printf("Received message: %+v\n", msg) // Логирование полученного сообщения
    if err := DB.Create(&msg).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

// Обработчик для получения всех сообщений
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	if err := DB.Find(&messages).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Отладка: выводим полученные сообщения
	fmt.Printf("Retrieved messages: %+v\n", messages)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// Обработчик для приветствия
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", message)
}

// Обработчик для обновления сообщения
func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req MessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	message = req.Message
	fmt.Fprintln(w, "Message updated successfully")
}
