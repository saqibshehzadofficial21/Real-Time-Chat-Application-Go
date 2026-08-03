package routes

import (
    "chat-app/internal/handlers"
    "chat-app/internal/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(authH *handlers.AuthHandler, userH *handlers.UserHandler, msgH *handlers.MessageHandler, friendH *handlers.FriendHandler) *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.POST("/register", userH.Register)
        api.POST("/login", authH.Login)

        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.POST("/messages", msgH.SendMessage)
            protected.GET("/conversations/:id/messages", msgH.GetMessages)

            protected.POST("/friend-requests", friendH.SendRequest)
            protected.POST("/friend-requests/:id/accept", friendH.AcceptRequest)
            protected.POST("/friend-requests/:id/reject", friendH.RejectRequest)
            protected.GET("/friend-requests/pending", friendH.ListPending)
        }
    }

    return router
}