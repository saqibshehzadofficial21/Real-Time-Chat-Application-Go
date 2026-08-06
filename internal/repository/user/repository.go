package user

import "gorm.io/gorm"

// userRepo UserRepository interface ka asal implementation hai
type userRepo struct {
    db *gorm.DB
}

// NewUserRepository ek naya UserRepository banata hai
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db: db}
}