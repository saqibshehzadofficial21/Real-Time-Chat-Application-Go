package websocket

// registerClient naye client ko is pod ke LOCAL memory map mein add karta hai,
// aur Redis Hash mein presence record bhi karta hai
func (h *Hub) registerClient(c *Client) {
	if h.Clients[c.ConversationID] == nil {
		h.Clients[c.ConversationID] = make(map[*Client]bool)
	}
	h.Clients[c.ConversationID][c] = true

	h.addPresence(c.ConversationID, c.UserID)
}