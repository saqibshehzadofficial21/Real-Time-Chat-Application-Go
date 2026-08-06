package auth

import "chat-app/internal/models"

// AuthService batata hai authentication se related kya operations mumkin hain
type AuthService interface {
    Login(email, password string) (string, *models.User, error)
}