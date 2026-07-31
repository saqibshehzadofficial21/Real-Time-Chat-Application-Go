package main

import (
    "log"
    "chat-app/internal/config"
    "chat-app/internal/db"
    "chat-app/internal/handlers"
    "chat-app/internal/repository"
    "chat-app/internal/service"
    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.LoadConfig()

    dbConn, err := db.ConnectDB(cfg)
    if err != nil {
        log.Fatal("Database connection failed: ", err)
    }
    defer dbConn.Close()
    log.Println("Database connected successfully")

    userRepo := repository.NewUserRepository(dbConn)
    userService := service.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    router := gin.Default()

    api := router.Group("/api")
    {
        api.POST("/register", userHandler.Register)
    }

    log.Println("Server running on :8080")
    router.Run(":8080")
}

