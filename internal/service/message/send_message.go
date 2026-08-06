package message

import (
    "errors"
    "strings"
    "chat-app/internal/models"
)

// SendMessage naya message bhejta hai — pehle check karta hai sender member hai ya nahi
func (s *messageService) SendMessage(convID, senderID int, content string) (*models.Message, error) {
    if strings.TrimSpace(content) == "" {
        return nil, errors.New("message content cannot be empty")
    }

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