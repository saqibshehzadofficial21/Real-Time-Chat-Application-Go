package user

import (
    "errors"
    "chat-app/internal/models"
    "golang.org/x/crypto/bcrypt"
)

// Register naya user banata hai — password hash karke database mein save karta hai
func (s *userService) Register(username, email, password string) (*models.User, error) {
    if len(password) < 6 {
        return nil, errors.New("password must be at least 6 characters")
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    u := &models.User{
        Username:     username,
        Email:        email,
        PasswordHash: string(hashed),
    }

    if err := s.repo.Create(u); err != nil {
        return nil, err
    }
    return u, nil
}