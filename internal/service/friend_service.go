package service

import (
    "errors"
    "chat-app/internal/models"
    "chat-app/internal/repository"
)

type FriendService interface {
    SendRequest(senderID, receiverID int) (*models.FriendRequest, error)
    AcceptRequest(requestID, userID int) (*models.Conversation, error)
    RejectRequest(requestID, userID int) error
    ListPendingRequests(userID int) ([]models.FriendRequest, error)
}

type friendService struct {
    repo repository.FriendRepository
}

func NewFriendService(repo repository.FriendRepository) FriendService {
    return &friendService{repo: repo}
}

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

func (s *friendService) AcceptRequest(requestID, userID int) (*models.Conversation, error) {
    req, err := s.repo.GetRequestByID(requestID)
    if err != nil {
        return nil, errors.New("friend request not found")
    }

    // Sirf woh user accept kar sakta hai jisay request bheji gayi thi
    if req.ReceiverID != userID {
        return nil, errors.New("you are not authorized to accept this request")
    }

    if req.Status != models.StatusPending {
        return nil, errors.New("this request has already been processed")
    }

    if err := s.repo.UpdateStatus(requestID, models.StatusAccepted); err != nil {
        return nil, err
    }

    // Accept hote hi ek-taraf 1-to-1 conversation ban jayegi
    conv, err := s.repo.CreateConversation(false)
    if err != nil {
        return nil, err
    }

    if err := s.repo.AddParticipants(conv.ID, []int{req.SenderID, req.ReceiverID}); err != nil {
        return nil, err
    }

    return conv, nil
}

func (s *friendService) RejectRequest(requestID, userID int) error {
    req, err := s.repo.GetRequestByID(requestID)
    if err != nil {
        return errors.New("friend request not found")
    }

    if req.ReceiverID != userID {
        return errors.New("you are not authorized to reject this request")
    }

    return s.repo.UpdateStatus(requestID, models.StatusRejected)
}

func (s *friendService) ListPendingRequests(userID int) ([]models.FriendRequest, error) {
    return s.repo.GetPendingForUser(userID)
}