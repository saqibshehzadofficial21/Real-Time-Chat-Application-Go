package service

import (
    "errors"
    "chat-app/internal/models"
    "chat-app/internal/repository"
    "golang.org/x/crypto/bcrypt"
)

type AuthService interface {
    Login(email, password string) (*models.User, error)
}

type authService struct {
    repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
    return &authService{repo: repo}
}

func (s *authService) Login(email, password string) (*models.User, error) {
    user, err := s.repo.GetByEmail(email)
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, errors.New("invalid email or password")
    }

    return user, nil
}