package friend

import "chat-app/internal/models"

// AddParticipants diye gaye users ko conversation ka member bana deta hai
func (r *friendRepo) AddParticipants(conversationID int, userIDs []int) error {
    var participants []models.ConversationParticipant
    for _, uid := range userIDs {
        participants = append(participants, models.ConversationParticipant{
            ConversationID: conversationID,
            UserID:         uid,
        })
    }
    return r.db.Create(&participants).Error
}