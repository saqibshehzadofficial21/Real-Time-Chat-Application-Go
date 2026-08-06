package friend

import "chat-app/internal/models"

// UpdateStatus friend request ka status badalta hai (pending/accepted/rejected)
func (r *friendRepo) UpdateStatus(id int, status models.RequestStatus) error {
    return r.db.Model(&models.FriendRequest{}).Where("id = ?", id).Update("status", status).Error
}