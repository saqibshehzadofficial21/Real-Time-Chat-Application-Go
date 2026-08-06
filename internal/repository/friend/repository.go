package friend

import "gorm.io/gorm"

type friendRepo struct {
    db *gorm.DB
}

// NewFriendRepository ek naya FriendRepository banata hai
func NewFriendRepository(db *gorm.DB) FriendRepository {
    return &friendRepo{db: db}
}