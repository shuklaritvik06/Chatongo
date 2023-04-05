package websockets

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var rooms = make(map[string]Room)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	room := NewRoom(mux.Vars(r)["id"])
	rooms[room.ID] = *room
	go room.Run()
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Created Room Successfully!",
	})
}

func RoomRouteHandler(w http.ResponseWriter, r *http.Request) {
	room := rooms[mux.Vars(r)["id"]]
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		room:    room,
		conn:    socket,
		recieve: make(chan *Message),
	}
	room.Register <- client
	defer func() { room.Unregsister <- client }()
	go client.write()
	go room.Run()
	client.read()
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
