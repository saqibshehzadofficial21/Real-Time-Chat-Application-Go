package friend

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// RejectRequest POST /api/friend-requests/:id/reject ko handle karta hai
func (h *FriendHandler) RejectRequest(c *gin.Context) {
    reqID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request id"})
        return
    }

    userID := c.GetInt("userID")

    if err := h.service.RejectRequest(reqID, userID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "friend request rejected"})
}