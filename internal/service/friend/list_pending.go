package friend

import "chat-app/internal/models"

// ListPendingRequests user ki saari pending requests deta hai
func (s *friendService) ListPendingRequests(userID int) ([]models.FriendRequest, error) {
    return s.repo.GetPendingForUser(userID)
}