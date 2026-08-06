package user

import userrepo "chat-app/internal/repository/user"

type userService struct {
    repo userrepo.UserRepository
}

// NewUserService ek naya UserService banata hai
func NewUserService(repo userrepo.UserRepository) UserService {
    return &userService{repo: repo}
}