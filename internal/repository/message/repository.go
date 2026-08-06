package message

import "gorm.io/gorm"

type messageRepo struct {
    db *gorm.DB
}

// NewMessageRepository ek naya MessageRepository banata hai
func NewMessageRepository(db *gorm.DB) MessageRepository {
    return &messageRepo{db: db}
}