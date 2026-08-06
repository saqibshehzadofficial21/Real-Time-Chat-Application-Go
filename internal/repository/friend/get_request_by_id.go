package friend

import "chat-app/internal/models"

// GetRequestByID ek friend request ko uski ID se dhoondta hai
func (r *friendRepo) GetRequestByID(id int) (*models.FriendRequest, error) {
    var req models.FriendRequest
    if err := r.db.First(&req, id).Error; err != nil {
        return nil, err
    }
    return &req, nil
}