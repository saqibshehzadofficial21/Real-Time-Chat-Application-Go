package websocket

import (
    "encoding/json"
    "log"
    myredis "chat-app/internal/redis"
)

// ListenToRedis Redis ke saare "chat:conversation:*" channels ko subscribe karta hai
// Jab bhi KOI pod (khud sameet) publish kare, yeh function chalta hai
func (h *Hub) ListenToRedis() {
    pubsub := h.RedisClient.PSubscribe(myredis.Ctx, "chat:conversation:*")
    defer pubsub.Close()

    ch := pubsub.Channel()

    for redisMsg := range ch {
        var msg WSMessage
        if err := json.Unmarshal([]byte(redisMsg.Payload), &msg); err != nil {
            log.Println("Redis message parse error:", err)
            continue
        }

        // Is pod ke LOCAL clients ko bhejo (agar koi connected ho is conversation mein)
        h.broadcastMessage(msg)
    }
}