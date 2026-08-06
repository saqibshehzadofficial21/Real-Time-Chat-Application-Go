package friend

import friendrepo "chat-app/internal/repository/friend"

type friendService struct {
    repo friendrepo.FriendRepository
}

// NewFriendService ek naya FriendService banata hai
func NewFriendService(repo friendrepo.FriendRepository) FriendService {
    return &friendService{repo: repo}
}