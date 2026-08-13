package redis


import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

// NewRedisClient Redis se connection banata hai
func NewRedisClient(host, port string) *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%s", host, port),
    })
    return client
}