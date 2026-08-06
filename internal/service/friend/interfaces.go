
package friend

import "chat-app/internal/models"

// FriendService batata hai friend requests se related kya operations mumkin hain
type FriendService interface {
    SendRequest(senderID, receiverID int) (*models.FriendRequest, error)
    AcceptRequest(requestID, userID int) (*models.Conversation, error)
    RejectRequest(requestID, userID int) error
    ListPendingRequests(userID int) ([]models.FriendRequest, error)
}