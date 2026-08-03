package models

import "time"

type Message struct {
    ID             int       `json:"id" gorm:"primaryKey"`
    ConversationID int       `json:"conversation_id" gorm:"not null"`
    SenderID       int       `json:"sender_id" gorm:"not null"`
    Content        string    `json:"content" gorm:"not null"`
    IsRead         bool      `json:"is_read" gorm:"default:false"`
    CreatedAt      time.Time `json:"created_at"`
}

func (Message) TableName() string {
    return "messages"
}