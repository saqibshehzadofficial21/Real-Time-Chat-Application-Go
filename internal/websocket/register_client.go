package websocket

// registerClient naye client ko uski conversation ke room mein add karta hai
func (h *Hub) registerClient(c *Client) {
    if h.Clients[c.ConversationID] == nil {
        h.Clients[c.ConversationID] = make(map[*Client]bool)
    }
    h.Clients[c.ConversationID][c] = true
}