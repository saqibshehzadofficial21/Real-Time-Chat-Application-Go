package user

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Register POST /api/register ko handle karta hai
func (h *UserHandler) Register(c *gin.Context) {
    var req registerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    u, err := h.service.Register(req.Username, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, u)
}