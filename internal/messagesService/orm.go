package messagesService

import (
	"gorm.io/gorm"
)

// Message структура для хранения сообщений
type Message struct {
    gorm.Model
    Text string `json:"message"` // Тег JSON должен совпадать с ключом в JSON
}
