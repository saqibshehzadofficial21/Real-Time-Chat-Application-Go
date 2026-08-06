package message

import msgrepo "chat-app/internal/repository/message"

type messageService struct {
    repo msgrepo.MessageRepository
}

// NewMessageService ek naya MessageService banata hai
func NewMessageService(repo msgrepo.MessageRepository) MessageService {
    return &messageService{repo: repo}
}