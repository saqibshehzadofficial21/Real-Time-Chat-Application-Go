package websocket

func (h *Hub) unregisterClient(c *Client) {
	if clients, ok := h.Clients[c.ConversationID]; ok {
		if _, ok := clients[c]; ok {
			delete(clients, c)
			close(c.Send)
		}
	}

	h.removePresence(c.ConversationID, c.UserID)
}