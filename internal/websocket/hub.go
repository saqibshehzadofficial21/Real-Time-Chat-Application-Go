package websocket

import "github.com/redis/go-redis/v9"

// Hub saare LOCAL connected clients ko track karta hai (is pod ke andar)
type Hub struct {
    Clients     map[int]map[*Client]bool
    Register    chan *Client
    Unregister  chan *Client
    Broadcast   chan WSMessage
    RedisClient *redis.Client
}

// NewHub ek naya Hub banata hai, Redis client ke sath
func NewHub(redisClient *redis.Client) *Hub {
    return &Hub{
        Clients:     make(map[int]map[*Client]bool),
        Register:    make(chan *Client),
        Unregister:  make(chan *Client),
        Broadcast:   make(chan WSMessage),
        RedisClient: redisClient,
    }
}