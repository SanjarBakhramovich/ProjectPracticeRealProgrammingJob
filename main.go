package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message = "global variable"

type MessageRequest struct {
	Message string `json:"message"`
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

// 
// Обработчик для создания сообщения
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	// Считывание JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Сохранение сообщения в базу данных
	if err := DB.Create(&msg).Error; err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}


// Обработчик для получения всех сообщений
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	// Извлекаем все сообщения из базы данных
	if err := DB.Find(&messages).Error; err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Устанавливаем заголовок ответа
	w.Header().Set("Content-Type", "application/json")	
	// Кодируем список сообщений в JSON и отправляем ответ
	json.NewEncoder(w).Encode(messages)
}



func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/update-message", UpdateMessageHandler).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}