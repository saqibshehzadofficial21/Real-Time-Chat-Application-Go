package websocket

// Run Hub ka main loop hai
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.registerClient(client)

        case client := <-h.Unregister:
            h.unregisterClient(client)

        case msg := <-h.Broadcast:
            // Sirf Redis pe publish karo — ListenToRedis khud is pod ke
            // local clients ko bhi deliver kar dega, chahe message kahin se aaya ho
            h.publishToRedis(msg)
        }
    }
}