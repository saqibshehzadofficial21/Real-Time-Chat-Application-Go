package auth

import userrepo "chat-app/internal/repository/user"

type authService struct {
    repo userrepo.UserRepository
}

// NewAuthService ek naya AuthService banata hai
func NewAuthService(repo userrepo.UserRepository) AuthService {
    return &authService{repo: repo}
}