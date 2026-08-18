package websocket

import (
	"fmt"
	"os"
	myredis "chat-app/internal/redis"
)

// podName is har pod ka unique identifier — Kubernetes khud yeh env variable deta hai
func podName() string {
	name := os.Getenv("HOSTNAME") // Kubernetes mein pod ka naam yahan hota hai
	if name == "" {
		name = "local"
	}
	return name
}

// addPresence Redis Hash mein record karta hai "yeh user, is conversation mein, is pod pe hai"
func (h *Hub) addPresence(conversationID, userID int) {
	key := fmt.Sprintf("presence:conversation:%d", conversationID)
	field := fmt.Sprintf("user:%d", userID)
	h.RedisClient.HSet(myredis.Ctx, key, field, podName())
}

// removePresence disconnect hone par Hash se record hata deta hai
func (h *Hub) removePresence(conversationID, userID int) {
	key := fmt.Sprintf("presence:conversation:%d", conversationID)
	field := fmt.Sprintf("user:%d", userID)
	h.RedisClient.HDel(myredis.Ctx, key, field)
}