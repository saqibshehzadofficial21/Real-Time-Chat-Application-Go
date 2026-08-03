package handlers

import (
    "net/http"
    "strconv"
    "chat-app/internal/service"
    "github.com/gin-gonic/gin"
)

type FriendHandler struct {
    service service.FriendService
}

func NewFriendHandler(s service.FriendService) *FriendHandler {
    return &FriendHandler{service: s}
}

type sendRequestBody struct {
    ReceiverID int `json:"receiver_id" binding:"required"`
}

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

    c.JSON(http.StatusOK, gin.H{
        "message":      "friend request accepted",
        "conversation": conv,
    })
}

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

func (h *FriendHandler) ListPending(c *gin.Context) {
    userID := c.GetInt("userID")

    requests, err := h.service.ListPendingRequests(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, requests)
}