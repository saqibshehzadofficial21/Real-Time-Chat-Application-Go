package websocket

import "github.com/gorilla/websocket"

// Client ek connected user ka WebSocket connection represent karta hai
type Client struct {
    Hub            *Hub
    Conn           *websocket.Conn
    Send           chan []byte
    UserID         int
    ConversationID int
}

// NewClient ek naya Client banata hai
func NewClient(hub *Hub, conn *websocket.Conn, userID, conversationID int) *Client {
    return &Client{
        Hub:            hub,
        Conn:           conn,
        Send:           make(chan []byte, 256),
        UserID:         userID,
        ConversationID: conversationID,
    }
}