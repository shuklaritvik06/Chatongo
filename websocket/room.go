package websockets

type Room struct {
	ID          string
	Clients     map[*Client]bool
	Broadcast   chan *Message
	Register    chan *Client
	Unregsister chan *Client
}

type Message struct {
	Message  string
	Type     string
	ClientID string
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.Clients[client] = true
		case client := <-r.Unregsister:
			if _, ok := r.Clients[client]; ok {
				delete(r.Clients, client)
				close(client.recieve)
			}
		case message := <-r.Broadcast:
			for client := range r.Clients {
				client.recieve <- message
			}
		}
	}
}
