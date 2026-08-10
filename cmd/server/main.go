package main

import (
    "log"

    "chat-app/internal/config"
    "chat-app/internal/db"
    "chat-app/internal/models"
    "chat-app/internal/routes"
    "chat-app/internal/websocket"
    "chat-app/pkg/utils"

    authH "chat-app/internal/handlers/auth"
    friendH "chat-app/internal/handlers/friend"
    groupH "chat-app/internal/handlers/group"
    messageH "chat-app/internal/handlers/message"
    userH "chat-app/internal/handlers/user"
    wsH "chat-app/internal/handlers/websocket"

    authSvc "chat-app/internal/service/auth"
    friendSvc "chat-app/internal/service/friend"
    groupSvc "chat-app/internal/service/group"
    messageSvc "chat-app/internal/service/message"
    userSvc "chat-app/internal/service/user"

    friendRepo "chat-app/internal/repository/friend"
    groupRepo "chat-app/internal/repository/group"
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
        &models.User{}, &models.Conversation{}, &models.ConversationParticipant{},
        &models.Message{}, &models.FriendRequest{},
    ); err != nil {
        log.Fatal("AutoMigrate failed: ", err)
    }
    log.Println("Database connected and migrated successfully")

    uRepo := userRepo.NewUserRepository(dbConn)
    mRepo := messageRepo.NewMessageRepository(dbConn)
    fRepo := friendRepo.NewFriendRepository(dbConn)
    gRepo := groupRepo.NewGroupRepository(dbConn)

    uService := userSvc.NewUserService(uRepo)
    aService := authSvc.NewAuthService(uRepo)
    mService := messageSvc.NewMessageService(mRepo)
    fService := friendSvc.NewFriendService(fRepo)
    gService := groupSvc.NewGroupService(gRepo)

    uHandler := userH.NewUserHandler(uService)
    aHandler := authH.NewAuthHandler(aService)
    mHandler := messageH.NewMessageHandler(mService)
    fHandler := friendH.NewFriendHandler(fService)
    gHandler := groupH.NewGroupHandler(gService)

    hub := websocket.NewHub()
    go hub.Run()
    wsHandler := wsH.NewWSHandler(hub, mService)

    router := routes.SetupRoutes(aHandler, uHandler, mHandler, fHandler, gHandler, wsHandler)

    log.Println("Server running on :8080")
    router.Run(":8080")
}