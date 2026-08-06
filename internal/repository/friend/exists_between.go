package friend

import "chat-app/internal/models"

// ExistsBetween check karta hai ke do users ke beech pehle se koi request maujood hai ya nahi
func (r *friendRepo) ExistsBetween(senderID, receiverID int) (bool, error) {
    var count int64
    err := r.db.Model(&models.FriendRequest{}).
        Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
            senderID, receiverID, receiverID, senderID).
        Count(&count).Error
    return count > 0, err
}