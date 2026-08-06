package friend

import (
    "errors"
    "chat-app/internal/models"
)

// AcceptRequest request accept karta hai aur khud-ba-khud conversation bana deta hai
func (s *friendService) AcceptRequest(requestID, userID int) (*models.Conversation, error) {
    req, err := s.repo.GetRequestByID(requestID)
    if err != nil {
        return nil, errors.New("friend request not found")
    }

    if req.ReceiverID != userID {
        return nil, errors.New("you are not authorized to accept this request")
    }
    if req.Status != models.StatusPending {
        return nil, errors.New("this request has already been processed")
    }

    if err := s.repo.UpdateStatus(requestID, models.StatusAccepted); err != nil {
        return nil, err
    }

    conv, err := s.repo.CreateConversation(false)
    if err != nil {
        return nil, err
    }

    if err := s.repo.AddParticipants(conv.ID, []int{req.SenderID, req.ReceiverID}); err != nil {
        return nil, err
    }

    return conv, nil
}