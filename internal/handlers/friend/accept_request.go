package friend

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// AcceptRequest POST /api/friend-requests/:id/accept ko handle karta hai
func (h *FriendHandler) AcceptRequest(c *gin.Context) {
    reqID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request id"})
        return
    }

    userID := c.GetInt("userID")

    conv, err := h.service.AcceptRequest(reqID, userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "friend request accepted", "conversation": conv})
}