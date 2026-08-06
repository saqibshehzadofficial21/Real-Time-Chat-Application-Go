package message

import (
    "errors"
    "chat-app/internal/models"
)

// GetConversationMessages conversation ke messages deta hai — sirf agar requester member ho
func (s *messageService) GetConversationMessages(convID, requesterID int) ([]models.Message, error) {
    isMember, err := s.repo.IsParticipant(convID, requesterID)
    if err != nil {
        return nil, err
    }
    if !isMember {
        return nil, errors.New("you are not a participant of this conversation")
    }

    return s.repo.GetByConversationID(convID)
}