package websocket

// WSMessage woh JSON format hai jo client aur server ke beech WebSocket par bhejte hain
type WSMessage struct {
    Type           string `json:"type"`            // "message" ya "typing"
    ConversationID int    `json:"conversation_id"`
    SenderID       int    `json:"sender_id"`
    Content        string `json:"content"`
}