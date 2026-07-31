package repository

import "database/sql"

type ConversationRepository interface {
    Create(isGroup bool) (int, error)
}

type conversationRepo struct {
    db *sql.DB
}

func NewConversationRepository(db *sql.DB) ConversationRepository {
    return &conversationRepo{db: db}
}

func (r *conversationRepo) Create(isGroup bool) (int, error) {
    var id int
    query := `INSERT INTO conversations (is_group) VALUES ($1) RETURNING id`
    err := r.db.QueryRow(query, isGroup).Scan(&id)
    return id, err
}