package friend

import (
    "errors"
    "chat-app/internal/models"
)

// SendRequest naya friend request bhejta hai — duplicate aur self-request check karta hai
func (s *friendService) SendRequest(senderID, receiverID int) (*models.FriendRequest, error) {
    if senderID == receiverID {
        return nil, errors.New("cannot send friend request to yourself")
    }

    exists, err := s.repo.ExistsBetween(senderID, receiverID)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("friend request already exists between these users")
    }

    req := &models.FriendRequest{
        SenderID:   senderID,
        ReceiverID: receiverID,
        Status:     models.StatusPending,
    }

    if err := s.repo.CreateRequest(req); err != nil {
        return nil, err
    }
    return req, nil
}