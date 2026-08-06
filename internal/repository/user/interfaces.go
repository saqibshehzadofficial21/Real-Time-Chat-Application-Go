package user

import "chat-app/internal/models"

// UserRepository batata hai user data ke sath kya operations mumkin hain
type UserRepository interface {
    Create(user *models.User) error
    GetByEmail(email string) (*models.User, error)
}
