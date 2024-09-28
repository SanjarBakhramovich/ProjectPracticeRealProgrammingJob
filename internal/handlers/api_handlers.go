package handlers

import (
	"REST/internal/messagesService"
	messages "REST/internal/messagesService"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Обработчики для API
type Handler struct {
	Service *messagesService.MessageService
}

// Конструктор для создания структуры Handler
func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

// GET
// Получить все сообщения
func (h *Handler) GetMessagesHandler(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()

	if err != nil {
		return nil, err
	}

	var response []messages.Message

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      msg.ID,
			Message: msg.Text,
		}
		response = append(response, message)
	}

	return messages.GetMessages200JSONResponse{
		JSON200: &response,
	}, nil
}

// POST
// Создать новое сообщение
func (h *Handler) PostMessageHandler(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body

	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
        Id:      &createdMessage.ID,
        Message: &createdMessage.Text,
    }
	return response, nil
}

// PATCH /api/patch/{id}
// Обработчик для обновления сообщения
func (h *Handler) PatchMessageHandler(c echo.Context) error{
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Недопустимый ID"})
	}

	var message messagesService.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedMessage, err := h.Service.UpdateMessageByID(id, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedMessage)
}

// DELETE /api/delete/{id}
// Обработчик для удаления сообщения
func (h *Handler) DeleteMessageHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Недопустимый ID"})
	}

	err = h.Service.DeleteMessageByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusOK)
}
