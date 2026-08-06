package message

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// GetMessages GET /api/conversations/:id/messages ko handle karta hai
func (h *MessageHandler) GetMessages(c *gin.Context) {
    convID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
        return
    }

    requesterID := c.GetInt("userID")

    messages, err := h.service.GetConversationMessages(convID, requesterID)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, messages)
}