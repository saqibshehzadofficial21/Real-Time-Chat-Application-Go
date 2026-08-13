package websocket

import (
    "encoding/json"
    "fmt"
    "log"
    myredis "chat-app/internal/redis"
)

// broadcastMessage LOCAL clients ko bhejta hai (is pod ke andar jo connected hain)
func (h *Hub) broadcastMessage(msg WSMessage) {
    data, err := json.Marshal(msg)
    if err != nil {
        return
    }

    for client := range h.Clients[msg.ConversationID] {
        select {
        case client.Send <- data:
        default:
            close(client.Send)
            delete(h.Clients[msg.ConversationID], client)
        }
    }
}

// publishToRedis message ko Redis channel pe publish karta hai — taake BAAKI pods bhi sun sakein
func (h *Hub) publishToRedis(msg WSMessage) {
    data, err := json.Marshal(msg)
    if err != nil {
        return
    }

    channel := fmt.Sprintf("chat:conversation:%d", msg.ConversationID)
    if err := h.RedisClient.Publish(myredis.Ctx, channel, data).Err(); err != nil {
        log.Println("Redis publish error:", err)
    }
}