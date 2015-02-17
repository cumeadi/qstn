package main

import "github.com/gorilla/websocket"

type Hub struct {
	clients map[*websocket.Conn]*Client
}

func NewHub() *Hub {
	return &Hub{map[*websocket.Conn]*Client{}}
}

func (h *Hub) Add(ws *websocket.Conn) *Client {
	h.clients[ws] = &Client{ws, make(chan struct{})}
	return h.clients[ws]
}

func (h *Hub) Broadcast(v interface{}) {
	for ws, _ := range h.clients {
		ws.WriteJSON(v)
	}
}

func (h *Hub) Remove(ws *websocket.Conn) {
	delete(h.clients, ws)
}
