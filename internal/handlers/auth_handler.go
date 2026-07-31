package handlers

import (
    "net/http"
    "chat-app/internal/service"
    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
    return &AuthHandler{service: s}
}

type loginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req loginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.Login(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}