package messagesService

// MessageService - Сервис для работы с сообщениями
type MessageService struct{
	repo MessageRepository
}

// NewService - Создание нового сервиса
// repo - репозиторий для работы с сообщениями
// return *MessageService - возвращает новый сервис
func NewService(repo MessageRepository) *MessageService{
	// Инициализация сервиса
	return &MessageService{repo: repo}
}

// Create Message - Создание сообщения
func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}


// Get Messages - Получение всех сообщений
func (s *MessageService) GetAllMessages(message Message) ([]Message, error) {
	return s.repo.GetAllMessages()
}
