package friend

import "chat-app/internal/models"

// GetPendingForUser ek user ki saari pending (abhi tak accept/reject na hui) requests deta hai
func (r *friendRepo) GetPendingForUser(userID int) ([]models.FriendRequest, error) {
    var requests []models.FriendRequest
    err := r.db.Where("receiver_id = ? AND status = ?", userID, models.StatusPending).Find(&requests).Error
    return requests, err
}