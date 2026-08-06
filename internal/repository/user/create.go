package user

import "chat-app/internal/models"

// Create naya user database mein insert karta hai
func (r *userRepo) Create(u *models.User) error {
    return r.db.Create(u).Error
}