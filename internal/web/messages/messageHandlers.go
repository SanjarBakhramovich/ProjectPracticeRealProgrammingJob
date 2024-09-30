package messages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Структура для хранения сообщений в памяти (в реальном проекте это будет база данных)
var messages = []Message{
    {Id: 1, Message: "Hello, World!"},
    {Id: 2, Message: "Another message"},
}

// Структура для сообщения
type Message struct {
    Id      int    `json:"id"`
    Message string `json:"message"`
}

// Обработчик GET-запроса для получения сообщений
func GetMessagesHandler(c echo.Context) error {
    // В данном примере возвращаются все сообщения
    if len(messages) == 0 {
        return c.JSON(http.StatusNotFound, "No messages found")
    }
    
    // Возвращаем первый элемент как пример
    response := GetMessages200JSONResponse(messages)
    return c.JSON(http.StatusOK, response)
}

// Обработчик POST-запроса для создания нового сообщения
func PostMessagesHandler(c echo.Context) error {
    var request PostMessagesRequestObject
    
    // Парсим тело запроса
    if err := c.Bind(&request); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request format")
    }
    
    // Проверяем, что поле Message не пустое
    if request.Body == nil || request.Body.Message == nil {
        return c.JSON(http.StatusBadRequest, "Message is required")
    }
    
    // Создаем новое сообщение
    newMessage := Message{
        Id:      len(messages) + 1,
        Message: *request.Body.Message,
    }
    
    // Добавляем новое сообщение в список
    messages = append(messages, newMessage)
    
    // Формируем ответ
    response := PostMessagesResponseObject{
        Id:      &newMessage.Id,
        Message: &newMessage.Message,
    }
    
    return c.JSON(http.StatusCreated, response)
}

// Структуры для GET-запросов и ответов
type GetMessagesRequestObject struct {
    // Здесь можно указать параметры запроса, если они нужны
}

type GetMessagesResponseObject struct {
    Id      int    `json:"id"`
    Message string `json:"message"`
}

// Структуры для POST-запросов и ответов
type PostMessagesRequestObject struct {
    Body *PostMessageRequestBody `json:"body"`
}

type PostMessageRequestBody struct {
    Message *string `json:"message"`
}

type PostMessagesResponseObject struct {
    Id      *int    `json:"id"`
    Message *string `json:"message"`
}

// Ответ на успешный GET-запрос
func GetMessages200JSONResponse(response []Message) GetMessagesResponseObject {
    return GetMessagesResponseObject{
        Id:      response[0].Id,
        Message: response[0].Message,
    }
}
