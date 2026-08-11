package websocket

import (
    "time"
    "github.com/gorilla/websocket"
)

const (
    pongWait   = 60 * time.Second
    pingPeriod = (pongWait * 9) / 10 // pongWait se thora kam
)

// WritePump client ko messages bhejta hai aur periodic ping bhi karta hai taake connection zinda rahe
func (c *Client) WritePump() {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        ticker.Stop()
        c.Conn.Close()
    }()

    for {
        select {
        case data, ok := <-c.Send:
            if !ok {
                c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
                return
            }

        case <-ticker.C:
            // Har 54 second mein ek "ping" bhejte hain taake connection zinda rahe
            if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}