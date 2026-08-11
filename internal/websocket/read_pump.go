package websocket

import (
    "encoding/json"
    "log"
    "time"
)

// ReadPump client se aane wale messages padhta hai, aur pong response bhi handle karta hai
func (c *Client) ReadPump(onMessage func(WSMessage)) {
    defer func() {
        c.Hub.Unregister <- c
        c.Conn.Close()
    }()

    c.Conn.SetReadDeadline(time.Now().Add(pongWait))
    c.Conn.SetPongHandler(func(string) error {
        c.Conn.SetReadDeadline(time.Now().Add(pongWait))
        return nil
    })

    for {
        _, data, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println("read error:", err)
            break
        }

        var msg WSMessage
        if err := json.Unmarshal(data, &msg); err != nil {
            continue
        }

        onMessage(msg)
    }
}