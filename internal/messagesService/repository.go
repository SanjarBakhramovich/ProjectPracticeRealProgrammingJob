package messagesService

import "gorm.io/gorm"

// MessageRepository - Интерфейс репозитория для работы с сообщениями
type MessageRepository interface {
    CreateMessage(message Message) (Message, error)
    GetAllMessages() ([]Message, error)
    UpdateMessageByID(id int, message Message) (Message, error)
    DeleteMessageByID(id int) error
}
	// MessageRepository - Репозиторий для работы с сообщениями
type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

	// CreateMessage - Создание сообщения
func (r *messageRepository) CreateMessage(message Message) (Message, error){
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}
// Get GetAllMessages - Получение всех сообщений
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}


// Patch Message - Обновление сообщения по ID
func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	var existingMessage Message

    // Находим сообщение по ID
	if err := r.db.First(&existingMessage).Error; err != nil {
		return Message{}, err
	}
    // Обновляем сообщение
	if err := r.db.Model(&existingMessage).Updates(message).Error; err != nil{
		return Message{}, err
	}
    return existingMessage, nil
}


// Delete Message - Удаление сообщения
func (r *messageRepository) DeleteMessageByID(id int) error {
    if err := r.db.Delete(&Message{}, id).Error; err != nil {
		return err
	}
	return nil
}