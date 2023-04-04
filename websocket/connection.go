package websockets

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (new_room *Room) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		room:    new_room,
		conn:    socket,
		recieve: make(chan *Message),
	}
	new_room.Register <- client
	defer func() { new_room.Unregsister <- client }()
	go client.write()
	client.read()
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "success",
			"room_id": new_room.ID,
		},
	)
}

func NewRoom(id string) *Room {
	room := &Room{
		ID:          id,
		Broadcast:   make(chan *Message),
		Clients:     make(map[*Client]bool),
		Register:    make(chan *Client),
		Unregsister: make(chan *Client),
	}
	return room
}
