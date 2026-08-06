

package main

import (
    "log"

    "chat-app/internal/config"
    "chat-app/internal/db"
    "chat-app/internal/models"
    "chat-app/internal/routes"
    "chat-app/pkg/utils"

    authH "chat-app/internal/handlers/auth"
    friendH "chat-app/internal/handlers/friend"
    messageH "chat-app/internal/handlers/message"
    userH "chat-app/internal/handlers/user"

    authSvc "chat-app/internal/service/auth"
    friendSvc "chat-app/internal/service/friend"
    messageSvc "chat-app/internal/service/message"
    userSvc "chat-app/internal/service/user"

    friendRepo "chat-app/internal/repository/friend"
    messageRepo "chat-app/internal/repository/message"
    userRepo "chat-app/internal/repository/user"
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

    // Repository layer
    uRepo := userRepo.NewUserRepository(dbConn)
    mRepo := messageRepo.NewMessageRepository(dbConn)
    fRepo := friendRepo.NewFriendRepository(dbConn)

    // Service layer
    uService := userSvc.NewUserService(uRepo)
    aService := authSvc.NewAuthService(uRepo)
    mService := messageSvc.NewMessageService(mRepo)
    fService := friendSvc.NewFriendService(fRepo)

    // Handler layer
    uHandler := userH.NewUserHandler(uService)
    aHandler := authH.NewAuthHandler(aService)
    mHandler := messageH.NewMessageHandler(mService)
    fHandler := friendH.NewFriendHandler(fService)

    router := routes.SetupRoutes(aHandler, uHandler, mHandler, fHandler)

    log.Println("Server running on :8080")
    router.Run(":8080")
}