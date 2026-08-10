package group

import groupsvc "chat-app/internal/service/group"

// GroupHandler HTTP requests ko GroupService tak forward karta hai
type GroupHandler struct {
    service groupsvc.GroupService
}

// NewGroupHandler ek naya GroupHandler banata hai
func NewGroupHandler(s groupsvc.GroupService) *GroupHandler {
    return &GroupHandler{service: s}
}