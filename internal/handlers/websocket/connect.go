package websocket

import (
    "log"
    "net/http"
    "strconv"

    "chat-app/internal/websocket"
    "chat-app/pkg/utils"

    "github.com/gin-gonic/gin"
    gorillaws "github.com/gorilla/websocket"
)

var upgrader = gorillaws.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *WSHandler) Connect(c *gin.Context) {
    token := c.Query("token")
    convIDStr := c.Query("conversation_id")

    userID, err := utils.ValidateJWT(token)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing token"})
        return
    }

    convID, err := strconv.Atoi(convIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation_id"})
        return
    }

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }

    client := websocket.NewClient(h.Hub, conn, userID, convID)
    h.Hub.Register <- client
    log.Printf("Registered client userID=%d for conversation=%d", userID, convID)

    go client.WritePump()
    go client.ReadPump(func(msg websocket.WSMessage) {
        log.Printf("Received message from userID=%d: %s", userID, msg.Content)

        msg.SenderID = userID
        msg.ConversationID = convID

        if _, err := h.MessageService.SendMessage(convID, userID, msg.Content); err != nil {
            log.Printf("SendMessage failed: %v", err)
            return
        }

        log.Printf("Pushing message to Hub.Broadcast channel")
        h.Hub.Broadcast <- msg
        log.Printf("Message pushed to Hub.Broadcast channel successfully")
    })
}