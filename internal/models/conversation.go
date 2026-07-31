package models

import "time"

type Conversation struct {
    ID        int       `json:"id"`
    IsGroup   bool      `json:"is_group"`
    CreatedAt time.Time `json:"created_at"`
}