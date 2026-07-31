package service

import (
    "errors"
    "chat-app/internal/models"
    "chat-app/internal/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserService interface {
    Register(username, email, password string) (*models.User, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) Register(username, email, password string) (*models.User, error) {
    if len(password) < 6 {
        return nil, errors.New("password must be at least 6 characters")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Username:     username,
        Email:        email,
        PasswordHash: string(hashedPassword),
    }

    if err := s.repo.Create(user); err != nil {
        return nil, err
    }
    return user, nil
}