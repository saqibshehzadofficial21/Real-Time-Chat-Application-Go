package websocket

// Hub saare connected clients ko track karta hai aur unke beech messages route karta hai
type Hub struct {
    Clients    map[int]map[*Client]bool // conversationID -> set of clients
    Register   chan *Client
    Unregister chan *Client
    Broadcast  chan WSMessage
}

// NewHub ek naya Hub banata hai
func NewHub() *Hub {
    return &Hub{
        Clients:    make(map[int]map[*Client]bool),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Broadcast:  make(chan WSMessage),
    }
}