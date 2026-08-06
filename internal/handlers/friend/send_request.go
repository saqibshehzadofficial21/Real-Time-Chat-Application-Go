package friend

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// SendRequest POST /api/friend-requests ko handle karta hai
func (h *FriendHandler) SendRequest(c *gin.Context) {
    var body sendRequestBody
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    senderID := c.GetInt("userID")

    req, err := h.service.SendRequest(senderID, body.ReceiverID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, req)
}