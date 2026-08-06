

package routes

import (
    authH "chat-app/internal/handlers/auth"
    friendH "chat-app/internal/handlers/friend"
    messageH "chat-app/internal/handlers/message"
    userH "chat-app/internal/handlers/user"
    "chat-app/internal/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(auth *authH.AuthHandler, user *userH.UserHandler, msg *messageH.MessageHandler, friend *friendH.FriendHandler) *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.POST("/register", user.Register)
        api.POST("/login", auth.Login)

        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.POST("/messages", msg.SendMessage)
            protected.GET("/conversations/:id/messages", msg.GetMessages)

            protected.POST("/friend-requests", friend.SendRequest)
            protected.POST("/friend-requests/:id/accept", friend.AcceptRequest)
            protected.POST("/friend-requests/:id/reject", friend.RejectRequest)
            protected.GET("/friend-requests/pending", friend.ListPending)
        }
    }

    return router
}