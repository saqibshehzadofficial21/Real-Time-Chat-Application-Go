package group

import "chat-app/internal/models"

// GroupService batata hai group chat banane se related kya operations mumkin hain
type GroupService interface {
    CreateGroup(name string, creatorID int, memberIDs []int) (*models.Conversation, error)
}