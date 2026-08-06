package message

import "chat-app/internal/models"

// MessageService batata hai messaging se related kya operations mumkin hain
type MessageService interface {
    SendMessage(convID, senderID int, content string) (*models.Message, error)
    GetConversationMessages(convID, requesterID int) ([]models.Message, error)
}
