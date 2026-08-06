package user

import usersvc "chat-app/internal/service/user"

// UserHandler HTTP requests ko UserService tak forward karta hai
type UserHandler struct {
    service usersvc.UserService
}

// NewUserHandler ek naya UserHandler banata hai
func NewUserHandler(s usersvc.UserService) *UserHandler {
    return &UserHandler{service: s}
}