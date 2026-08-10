package websocket

import (
    "encoding/json"
    "log"
)

func (h *Hub) broadcastMessage(msg WSMessage) {
    data, err := json.Marshal(msg)
    if err != nil {
        return
    }

    log.Printf("Broadcasting to conversation %d — total clients found: %d", msg.ConversationID, len(h.Clients[msg.ConversationID]))

    for client := range h.Clients[msg.ConversationID] {
        select {
        case client.Send <- data:
            log.Printf("Sent to client userID=%d", client.UserID)
        default:
            close(client.Send)
            delete(h.Clients[msg.ConversationID], client)
        }
    }
}