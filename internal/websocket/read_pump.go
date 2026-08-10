package websocket

import (
    "encoding/json"
    "log"
)

// ReadPump client se aane wale messages padhta hai (yeh ek goroutine mein chalta hai)
func (c *Client) ReadPump(onMessage func(WSMessage)) {
    defer func() {
        c.Hub.Unregister <- c
        c.Conn.Close()
    }()

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