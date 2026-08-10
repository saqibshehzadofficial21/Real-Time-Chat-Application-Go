package websocket

// Run Hub ka main loop hai — yeh hamesha ek alag goroutine mein chalta hai
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.registerClient(client)

        case client := <-h.Unregister:
            h.unregisterClient(client)

        case msg := <-h.Broadcast:
            h.broadcastMessage(msg)
        }
    }
}