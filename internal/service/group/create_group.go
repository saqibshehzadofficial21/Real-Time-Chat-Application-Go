package group

import (
    "errors"
    "chat-app/internal/models"
)

// CreateGroup naya group banata hai — kam se kam 1 member (creator ke ilawa) zaroori hai
func (s *groupService) CreateGroup(name string, creatorID int, memberIDs []int) (*models.Conversation, error) {
    if name == "" {
        return nil, errors.New("group name is required")
    }
    if len(memberIDs) < 1 {
        return nil, errors.New("group must have at least one other member")
    }

    conv, err := s.repo.CreateGroup(name)
    if err != nil {
        return nil, err
    }

    allMembers := append(memberIDs, creatorID)

    if err := s.repo.AddMembers(conv.ID, allMembers); err != nil {
        return nil, errors.New("failed to add members to group: " + err.Error())
    }

    return conv, nil
}