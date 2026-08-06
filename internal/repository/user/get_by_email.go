package user

import "chat-app/internal/models"

// GetByEmail email se user dhoondta hai (login ke liye use hota hai)
func (r *userRepo) GetByEmail(email string) (*models.User, error) {
    var u models.User
    if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
        return nil, err
    }
    return &u, nil
}