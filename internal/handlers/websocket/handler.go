package websocket

import (
    "chat-app/internal/websocket"
    messagesvc "chat-app/internal/service/message"
)

// WSHandler HTTP connection ko WebSocket mein upgrade karta hai
type WSHandler struct {
    Hub            *websocket.Hub
    MessageService messagesvc.MessageService
}

// NewWSHandler ek naya WSHandler banata hai
func NewWSHandler(hub *websocket.Hub, msgService messagesvc.MessageService) *WSHandler {
    return &WSHandler{Hub: hub, MessageService: msgService}
}