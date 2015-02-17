package main

import "github.com/gorilla/websocket"
import "time"

type Client struct {
	ws   *websocket.Conn
	ping chan struct{}
}

func (c *Client) Ping(d time.Duration) {
	t := time.NewTicker(d)

	ping := struct {
		Ping bool `json:"ping"`
	}{
		true,
	}

	for {
		select {
		case <-t.C:
			c.ws.WriteJSON(ping)
		case <-c.ping:
			t.Stop()
			return
		}
	}
}
