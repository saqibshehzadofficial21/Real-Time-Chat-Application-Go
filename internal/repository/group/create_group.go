package group

import "chat-app/internal/models"

// CreateGroup naya group conversation banata hai (is_group = true)
func (r *groupRepo) CreateGroup(name string) (*models.Conversation, error) {
    conv := &models.Conversation{
        Name:    name,
        IsGroup: true,
    }
    err := r.db.Create(conv).Error
    return conv, err
}