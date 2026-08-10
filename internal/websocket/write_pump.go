package websocket

import "github.com/gorilla/websocket"

// WritePump client ko messages bhejta hai (yeh ek goroutine mein chalta hai)
func (c *Client) WritePump() {
    defer c.Conn.Close()

    for data := range c.Send {
        if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
            break
        }
    }
}