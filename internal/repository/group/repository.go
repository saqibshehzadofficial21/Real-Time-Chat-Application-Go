package group

import "gorm.io/gorm"

type groupRepo struct {
    db *gorm.DB
}

// NewGroupRepository ek naya GroupRepository banata hai
func NewGroupRepository(db *gorm.DB) GroupRepository {
    return &groupRepo{db: db}
}