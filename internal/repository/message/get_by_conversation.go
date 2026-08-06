package message

import "chat-app/internal/models"

// GetByConversationID ek conversation ke saare messages purane se naye order mein deta hai
func (r *messageRepo) GetByConversationID(convID int) ([]models.Message, error) {
    var messages []models.Message
    err := r.db.Where("conversation_id = ?", convID).
        Order("created_at asc").
        Find(&messages).Error
    return messages, err
}