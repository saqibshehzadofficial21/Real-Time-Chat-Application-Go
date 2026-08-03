package middleware

import (
    "time"
    "log"
    "github.com/gin-gonic/gin"
)

// LoggerMiddleware har request ka time aur path console mein print karta hai
func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        c.Next() // pehle asal handler chalne dein

        duration := time.Since(start)
        log.Printf("[%s] %s - %v", c.Request.Method, c.Request.URL.Path, duration)
    }
}