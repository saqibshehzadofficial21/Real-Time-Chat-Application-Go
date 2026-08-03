package repository

import (
    "chat-app/internal/models"
    "gorm.io/gorm"
)

type MessageRepository interface {
    Create(msg *models.Message) error
    GetByConversationID(convID int) ([]models.Message, error)
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