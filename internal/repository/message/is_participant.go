package message

import "chat-app/internal/models"

// IsParticipant check karta hai ke yeh user is conversation ka member hai ya nahi
func (r *messageRepo) IsParticipant(conversationID, userID int) (bool, error) {
    var count int64
    err := r.db.Model(&models.ConversationParticipant{}).
        Where("conversation_id = ? AND user_id = ?", conversationID, userID).
        Count(&count).Error
    return count > 0, err
}