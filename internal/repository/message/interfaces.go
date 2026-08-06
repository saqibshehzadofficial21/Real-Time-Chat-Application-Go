package message

import "chat-app/internal/models"

// MessageRepository batata hai message data ke sath kya operations mumkin hain
type MessageRepository interface {
    Create(msg *models.Message) error
    GetByConversationID(convID int) ([]models.Message, error)
    IsParticipant(conversationID, userID int) (bool, error)
}