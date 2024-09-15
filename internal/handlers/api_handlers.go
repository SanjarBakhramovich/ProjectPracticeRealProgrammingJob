package handlers

import (
	"REST/internal/messagesService"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	messages, err := h.Service.GetAllMessages() //исправил ошибку в задании
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

// PATCH /api/patch/{id}
// Обработчик для обновления сообщения
func (h *Handler) PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var message messagesService.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMessage, err := h.Service.UpdateMessageByID(id, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMessage)
}

// DELETE /api/delete/{id}
// Обработчик для удаления сообщения
func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteMessageByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
