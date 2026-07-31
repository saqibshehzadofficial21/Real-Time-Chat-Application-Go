package repository

import (
    "database/sql"
    "chat-app/internal/models"
)

type MessageRepository interface {
    Create(msg *models.Message) error
    GetByConversationID(convID int) ([]models.Message, error)
}

type messageRepo struct {
    db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
    return &messageRepo{db: db}
}

func (r *messageRepo) Create(msg *models.Message) error {
    query := `INSERT INTO messages (conversation_id, sender_id, content) 
              VALUES ($1, $2, $3) RETURNING id, created_at`
    return r.db.QueryRow(query, msg.ConversationID, msg.SenderID, msg.Content).
        Scan(&msg.ID, &msg.CreatedAt)
}

func (r *messageRepo) GetByConversationID(convID int) ([]models.Message, error) {
    rows, err := r.db.Query(`SELECT id, conversation_id, sender_id, content, is_read, created_at 
                              FROM messages WHERE conversation_id=$1 ORDER BY created_at ASC`, convID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var messages []models.Message
    for rows.Next() {
        var m models.Message
        if err := rows.Scan(&m.ID, &m.ConversationID, &m.SenderID, &m.Content, &m.IsRead, &m.CreatedAt); err != nil {
            return nil, err
        }
        messages = append(messages, m)
    }
    return messages, nil
}