
package websocket

// unregisterClient client ko disconnect hone par hata deta hai
func (h *Hub) unregisterClient(c *Client) {
    if clients, ok := h.Clients[c.ConversationID]; ok {
        if _, ok := clients[c]; ok {
            delete(clients, c)
            close(c.Send)
        }
    }
}