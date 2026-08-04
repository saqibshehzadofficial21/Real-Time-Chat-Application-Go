package repository

import (
    "chat-app/internal/models"
    "gorm.io/gorm"
)

type MessageRepository interface {
    Create(msg *models.Message) error
    GetByConversationID(convID int) ([]models.Message, error)
    IsParticipant(conversationID, userID int) (bool, error)
}

type messageRepo struct {
    db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
    return &messageRepo{db: db}
}

func (r *messageRepo) Create(msg *models.Message) error {
    return r.db.Create(msg).Error
}

func (r *messageRepo) GetByConversationID(convID int) ([]models.Message, error) {
    var messages []models.Message
    err := r.db.Where("conversation_id = ?", convID).
        Order("created_at asc").
        Find(&messages).Error
    return messages, err
}

// IsParticipant check karta hai ke yeh user is conversation ka member hai ya nahi
func (r *messageRepo) IsParticipant(conversationID, userID int) (bool, error) {
    var count int64
    err := r.db.Model(&models.ConversationParticipant{}).
        Where("conversation_id = ? AND user_id = ?", conversationID, userID).
        Count(&count).Error
    return count > 0, err
}