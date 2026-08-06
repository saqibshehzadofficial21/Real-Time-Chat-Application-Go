package message

import "chat-app/internal/models"

// Create naya message database mein insert karta hai
func (r *messageRepo) Create(msg *models.Message) error {
    return r.db.Create(msg).Error
}