package message

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// SendMessage POST /api/messages ko handle karta hai
func (h *MessageHandler) SendMessage(c *gin.Context) {
    var req sendMessageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    senderID := c.GetInt("userID")

    msg, err := h.service.SendMessage(req.ConversationID, senderID, req.Content)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, msg)
}