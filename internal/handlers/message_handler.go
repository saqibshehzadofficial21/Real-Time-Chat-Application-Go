package handlers

import (
    "net/http"
    "strconv"
    "chat-app/internal/service"
    "github.com/gin-gonic/gin"
)

type MessageHandler struct {
    service service.MessageService
}

func NewMessageHandler(s service.MessageService) *MessageHandler {
    return &MessageHandler{service: s}
}

type sendMessageRequest struct {
    ConversationID int    `json:"conversation_id" binding:"required"`
    Content        string `json:"content" binding:"required"`
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
    var req sendMessageRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    senderID := c.GetInt("userID") // middleware se aayega baad mein

    msg, err := h.service.SendMessage(req.ConversationID, senderID, req.Content)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, msg)
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
    convIDStr := c.Param("id")
    convID, err := strconv.Atoi(convIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
        return
    }

    messages, err := h.service.GetConversationMessages(convID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, messages)
}