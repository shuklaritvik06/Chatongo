package websockets

type Room struct {
	ID          string
	Clients     map[*Client]bool
	Broadcast   chan *Message
	Register    chan *Client
	Unregsister chan *Client
	Rooms       map[*Room]bool
}

type Message struct {
	Message  string
	Type     string
	ClientID string
}

func (r *Room) RegisterClient(client *Client) {
	message := &Message{
		Message: "New client joined the room",
		Type:    "notification",
	}
	r.BroadcastMessage(message)
	r.Clients[client] = true
}

func (r *Room) BroadcastMessage(message *Message) {
	for client := range r.Clients {
		client.recieve <- message
	}
}

func (r *Room) UnregisterClient(client *Client) {
	message := &Message{
		Message: "Client left the room",
		Type:    "notification",
	}
	r.BroadcastMessage(message)
	delete(r.Clients, client)
	close(client.recieve)
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.RegisterClient(client)
		case client := <-r.Unregsister:
			r.UnregisterClient(client)
		case message := <-r.Broadcast:
			r.BroadcastMessage(message)
		}
	}
}
