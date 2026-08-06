

package friend

import friendsvc "chat-app/internal/service/friend"

// FriendHandler HTTP requests ko FriendService tak forward karta hai
type FriendHandler struct {
    service friendsvc.FriendService
}

// NewFriendHandler ek naya FriendHandler banata hai
func NewFriendHandler(s friendsvc.FriendService) *FriendHandler {
    return &FriendHandler{service: s}
}