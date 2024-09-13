package handlers

import (
	"REST/internal/messagesService"
	"encoding/json"
	"net/http"
)

// Обработчики для API
type Handler struct{
	Service *messagesService.MessageService
}

// Конструктор для создания структуры Handler
func NewHandler(service *messagesService.MessageService) *Handler{
	return &Handler{
		Service: service,
	}
}

// GET
// Обработчик для получения всех сообщений
func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages(messagesService.Message{}) //исправил ошибку в задании
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// POST
// Обработчик для добавления нового сообщения
func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}


// // POST
// // Обработчик для создания сообщения
// func CreateMessage(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		Message string `json:"message"`
// 	}

//     if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }
	
//     fmt.Printf("Received message: %+v\n", input.Message) // Логирование полученного сообщения

// 	msg:= Message {
// 		Text : input.Message,
// 	}
//     if err := DB.Create(&msg).Error; err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
// 	w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(struct {
// 		ID uint `json:"id"`
// 		Message string `json:"message"`
// 	}{
// 		ID: msg.ID,
// 		Message: msg.Text,
// 	})
// }






// // PATCH
// // Обработчик для обновления сообщения
// func UpdateMessage(w http.ResponseWriter, r *http.Request) {
// 	// Получаем ID из URL
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	// Ищем сообщение по ID
// 	var msg Message
// 	if err := DB.First(&msg, id).Error; err != nil{
// 		http.Error(w, "Message not found", http.StatusNotFound)
// 		return
// 	}
	
// 	// Декодируем тело запроса
// 	var input struct {
// 		Message string `json:"message"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	// Обновляем сообщение
// 	msg.Text = input.Message
// 	if err := DB.Save(&msg).Error; err != nil {
// 		http.Error(w, "Failed to update message", http.StatusInternalServerError)
// 		return
// 	}

// 	// Возвращаем обновленное сообщение
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(struct {
// 		ID uint `json:"id"`
// 		Message string `json:"message"`
// 	}{
// 		ID: msg.ID,
// 		Message: msg.Text,
// 	})
// }

// // DELETE
// // Обработчик для удаления сообщения по ID
// func DeleteMessage(w http.ResponseWriter, r *http.Request) {	
// 	// Получаем ID из URL
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	// Ищем сообщение по ID
// 	var msg Message
// 	if err := DB.First(&msg, id).Error; err != nil {
// 		http.Error(w, "Message not found", http.StatusNotFound)
// 		return
// 	}

// 	// Удаляем сообщение
// 	if err := DB.Delete(&msg).Error; err != nil {
// 		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
// 		return
// 	}
	
// 	// Отправляем статус успешного удаления
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(struct {
// 		Message string `json:"message"`
// 	}{
// 		Message: "Message deleted successfully",
// 	})
// }

