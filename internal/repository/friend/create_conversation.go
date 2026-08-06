package friend

import "chat-app/internal/models"

// CreateConversation naya conversation banata hai (request accept hone par call hota hai)
func (r *friendRepo) CreateConversation(isGroup bool) (*models.Conversation, error) {
    conv := &models.Conversation{IsGroup: isGroup}
    err := r.db.Create(conv).Error
    return conv, err
}