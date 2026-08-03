package models

import "time"

type User struct {
    ID           int       `json:"id" gorm:"primaryKey"`
    Username     string    `json:"username" gorm:"unique;not null"`
    Email        string    `json:"email" gorm:"unique;not null"`
    PasswordHash string    `json:"-" gorm:"not null"`
    CreatedAt    time.Time `json:"created_at"`
}

// TableName GORM ko batata hai table ka exact naam kya hai
func (User) TableName() string {
    return "users"
}