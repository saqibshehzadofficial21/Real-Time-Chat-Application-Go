package group

import grouprepo "chat-app/internal/repository/group"

type groupService struct {
    repo grouprepo.GroupRepository
}

// NewGroupService ek naya GroupService banata hai
func NewGroupService(repo grouprepo.GroupRepository) GroupService {
    return &groupService{repo: repo}
}