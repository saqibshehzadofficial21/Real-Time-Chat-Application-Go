package group

import "chat-app/internal/models"

// GroupRepository batata hai group conversations ke sath kya operations mumkin hain
type GroupRepository interface {
    CreateGroup(name string) (*models.Conversation, error)
    AddMembers(conversationID int, userIDs []int) error
}
