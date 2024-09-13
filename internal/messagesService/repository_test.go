package messagesService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) CreateMessage(message Message) (Message, error) {
	args := m.Called(message)
	return args.Get(0).(Message), args.Error(1)
}

func (m *MockMessageRepository) GetAllMessages() ([]Message, error) {
	args := m.Called()
	return args.Get(0).([]Message), args.Error(1)
}

func (m *MockMessageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	args := m.Called(id, message)
	return args.Get(0).(Message), args.Error(1)
}

func (m *MockMessageRepository) DeleteMessageByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateMessage(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	message := Message{Model: gorm.Model{ID: 1}, Text: "Тестовое сообщение"} // Исправляем поля

	mockRepo.On("CreateMessage", message).Return(message, nil)

	createdMessage, err := mockRepo.CreateMessage(message)

	assert.NoError(t, err)
	assert.Equal(t, message, createdMessage)
	mockRepo.AssertExpectations(t)
}

func TestGetAllMessages(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	messages := []Message{
		{Model: gorm.Model{ID: 1}, Text: "Сообщение 1"},
		{Model: gorm.Model{ID: 2}, Text: "Сообщение 2"},
	}

	mockRepo.On("GetAllMessages").Return(messages, nil)

	retrievedMessages, err := mockRepo.GetAllMessages()

	assert.NoError(t, err)
	assert.Equal(t, messages, retrievedMessages)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMessageByID(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	message := Message{Model: gorm.Model{ID: 1}, Text: "Обновленное сообщение"}

	mockRepo.On("UpdateMessageByID", 1, message).Return(message, nil)

	updatedMessage, err := mockRepo.UpdateMessageByID(1, message)

	assert.NoError(t, err)
	assert.Equal(t, message, updatedMessage)
	mockRepo.AssertExpectations(t)
}

func TestDeleteMessageByID(t *testing.T) {
	mockRepo := new(MockMessageRepository)

	mockRepo.On("DeleteMessageByID", 1).Return(nil)

	err := mockRepo.DeleteMessageByID(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateMessageError(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	message := Message{Model: gorm.Model{ID: 1}, Text: "Тестовое сообщение"}
	expectedError := errors.New("ошибка создания")

	mockRepo.On("CreateMessage", message).Return(Message{}, expectedError)

	_, err := mockRepo.CreateMessage(message)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllMessagesError(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	expectedError := errors.New("ошибка получения")

	mockRepo.On("GetAllMessages").Return([]Message{}, expectedError)

	_, err := mockRepo.GetAllMessages()

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
