package auth

import (
    "errors"
    "chat-app/internal/models"
    "chat-app/pkg/utils"
    "golang.org/x/crypto/bcrypt"
)

// Login email/password verify karta hai aur success par JWT token deta hai
func (s *authService) Login(email, password string) (string, *models.User, error) {
    user, err := s.repo.GetByEmail(email)
    if err != nil {
        return "", nil, errors.New("invalid email or password")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", nil, errors.New("invalid email or password")
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return "", nil, errors.New("failed to generate token")
    }

    return token, user, nil
}