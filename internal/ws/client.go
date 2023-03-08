package ws

import "golang.org/x/net/websocket"

type Client struct {
	Conn     *websocket.Conn `json:"conn"`
	Message  chan *Message   `json:"message"`
	ID       string          `json:"id"`
	RoomID   string          `json:"room_id"`
	Username string          `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}
