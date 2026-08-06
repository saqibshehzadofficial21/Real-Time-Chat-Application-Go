package user

import "chat-app/internal/models"

// UserService batata hai user registration se related kya operations mumkin hain
type UserService interface {
    Register(username, email, password string) (*models.User, error)
}