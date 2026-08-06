package friend

import "chat-app/internal/models"

// CreateRequest nayi friend request database mein insert karta hai
func (r *friendRepo) CreateRequest(req *models.FriendRequest) error {
    return r.db.Create(req).Error
}