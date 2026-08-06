package friend

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ListPending GET /api/friend-requests/pending ko handle karta hai
func (h *FriendHandler) ListPending(c *gin.Context) {
    userID := c.GetInt("userID")

    requests, err := h.service.ListPendingRequests(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, requests)
}