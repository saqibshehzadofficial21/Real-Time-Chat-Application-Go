package friend

import "chat-app/internal/models"

// FriendRepository batata hai friend requests aur conversations ke sath kya operations mumkin hain
type FriendRepository interface {
    CreateRequest(req *models.FriendRequest) error
    GetRequestByID(id int) (*models.FriendRequest, error)
    UpdateStatus(id int, status models.RequestStatus) error
    GetPendingForUser(userID int) ([]models.FriendRequest, error)
    ExistsBetween(senderID, receiverID int) (bool, error)
    CreateConversation(isGroup bool) (*models.Conversation, error)
    AddParticipants(conversationID int, userIDs []int) error
}