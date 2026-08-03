package models

import "time"

type Conversation struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    IsGroup   bool      `json:"is_group" gorm:"default:false"`
    CreatedAt time.Time `json:"created_at"`
}

func (Conversation) TableName() string {
    return "conversations"
}