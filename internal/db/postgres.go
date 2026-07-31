package db

import (
    "database/sql"
    "fmt"
    "chat-app/internal/config"
    _ "github.com/lib/pq" // underscore import — driver register karne ke liye, direct use nahi hota
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )

    conn, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err = conn.Ping(); err != nil {
        return nil, err
    }

    return conn, nil
}