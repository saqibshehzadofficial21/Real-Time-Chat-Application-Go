package group

import "chat-app/internal/models"

// AddMembers diye gaye saare users ko group ka participant bana deta hai
func (r *groupRepo) AddMembers(conversationID int, userIDs []int) error {
    var participants []models.ConversationParticipant
    for _, uid := range userIDs {
        participants = append(participants, models.ConversationParticipant{
            ConversationID: conversationID,
            UserID:         uid,
        })
    }
    return r.db.Create(&participants).Error
}