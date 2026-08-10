package group

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// CreateGroup POST /api/groups ko handle karta hai
func (h *GroupHandler) CreateGroup(c *gin.Context) {
    var req createGroupRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    creatorID := c.GetInt("userID")

    conv, err := h.service.CreateGroup(req.Name, creatorID, req.MemberIDs)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, conv)
}