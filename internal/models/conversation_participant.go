package models

type ConversationParticipant struct {
    ConversationID int `json:"conversation_id" gorm:"primaryKey"`
    UserID         int `json:"user_id" gorm:"primaryKey"`
}

func (ConversationParticipant) TableName() string {
    return "conversation_participants"
}