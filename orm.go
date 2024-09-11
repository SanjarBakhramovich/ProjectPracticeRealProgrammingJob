package main

import (
	"gorm.io/gorm"
)

// Message структура для хранения сообщений
type Message struct {
    gorm.Model
    Text string `json:"text"` // Наш сервер будет ожидать JSON с полем text
}
