package main

import (
    "log"
    "chat-app/internal/config"
    "chat-app/internal/db"
    "chat-app/internal/handlers"
    "chat-app/internal/models"
    "chat-app/internal/repository"
    "chat-app/internal/routes"
    "chat-app/internal/service"
    "chat-app/pkg/utils"
)

func main() {
    cfg := config.LoadConfig()
    utils.InitJWTSecret(cfg.JWTSecret)

    dbConn, err := db.ConnectDB(cfg)
    if err != nil {
        log.Fatal("Database connection failed: ", err)
    }

    if err := dbConn.AutoMigrate(
        &models.User{},
        &models.Conversation{},
        &models.ConversationParticipant{},
        &models.Message{},
        &models.FriendRequest{},
    ); err != nil {
        log.Fatal("AutoMigrate failed: ", err)
    }
    log.Println("Database connected and migrated successfully")

    userRepo := repository.NewUserRepository(dbConn)
    msgRepo := repository.NewMessageRepository(dbConn)
    friendRepo := repository.NewFriendRepository(dbConn)

    userService := service.NewUserService(userRepo)
    authService := service.NewAuthService(userRepo)
    msgService := service.NewMessageService(msgRepo)
    friendService := service.NewFriendService(friendRepo)

    userHandler := handlers.NewUserHandler(userService)
    authHandler := handlers.NewAuthHandler(authService)
    msgHandler := handlers.NewMessageHandler(msgService)
    friendHandler := handlers.NewFriendHandler(friendService)

    router := routes.SetupRoutes(authHandler, userHandler, msgHandler, friendHandler)

    log.Println("Server running on :8080")
    router.Run(":8080")
}