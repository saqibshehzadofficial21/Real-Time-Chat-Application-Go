package repository

import (
    "database/sql"
    "chat-app/internal/models"
)

type UserRepository interface {
    Create(user *models.User) error
    GetByEmail(email string) (*models.User, error)
}

type userRepo struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
    query := `INSERT INTO users (username, email, password_hash) 
              VALUES ($1, $2, $3) RETURNING id, created_at`

    return r.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).
        Scan(&user.ID, &user.CreatedAt)
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
    user := &models.User{}
    query := `SELECT id, username, email, password_hash, created_at 
              FROM users WHERE email=$1`

    err := r.db.QueryRow(query, email).
        Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}