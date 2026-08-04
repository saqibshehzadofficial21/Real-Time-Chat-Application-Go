package service

import (
    "errors"
    "strings"
    "chat-app/internal/models"
    "chat-app/internal/repository"
)

type MessageService interface {
    SendMessage(convID, senderID int, content string) (*models.Message, error)
    GetConversationMessages(convID, requesterID int) ([]models.Message, error)
}

type messageService struct {
    repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
    return &messageService{repo: repo}
}

func (s *messageService) SendMessage(convID, senderID int, content string) (*models.Message, error) {
    if strings.TrimSpace(content) == "" {
        return nil, errors.New("message content cannot be empty")
    }

    // Authorization check: kya sender is conversation ka member hai?
    isMember, err := s.repo.IsParticipant(convID, senderID)
    if err != nil {
        return nil, err
    }
    if !isMember {
        return nil, errors.New("you are not a participant of this conversation")
    }

    msg := &models.Message{
        ConversationID: convID,
        SenderID:       senderID,
        Content:        content,
    }

    if err := s.repo.Create(msg); err != nil {
        return nil, err
    }
    return msg, nil
}

func (s *messageService) GetConversationMessages(convID, requesterID int) ([]models.Message, error) {
    // Authorization check: kya requester is conversation ka member hai?
    isMember, err := s.repo.IsParticipant(convID, requesterID)
    if err != nil {
        return nil, err
    }
    if !isMember {
        return nil, errors.New("you are not a participant of this conversation")
    }

    return s.repo.GetByConversationID(convID)
}