package repository

import (
    "chat-app/internal/models"
    "gorm.io/gorm"
)

type ConversationRepository interface {
    Create(isGroup bool) (*models.Conversation, error)
}

type conversationRepo struct {
    db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepository {
    return &conversationRepo{db: db}
}

func (r *conversationRepo) Create(isGroup bool) (*models.Conversation, error) {
    conv := &models.Conversation{IsGroup: isGroup}
    err := r.db.Create(conv).Error
    return conv, err
}