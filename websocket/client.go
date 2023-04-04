package websockets

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/shuklaritvik06/chatongo/backend/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Client struct {
	room    *Room
	conn    *websocket.Conn
	recieve chan *Message
}

func (c *Client) read() {
	defer c.conn.Close()
	for {
		messageType, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		c.room.Broadcast <- &Message{
			Message: string(msg),
			Type:    string(rune(messageType)),
		}
		database.GetDB().Database("chats").Collection(c.room.ID).InsertOne(context.Background(), bson.D{{
			Key: "message", Value: string(msg),
		}})
	}
}

func (c *Client) write() {
	defer c.conn.Close()
	for msg := range c.recieve {
		err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg.Message))
		if err != nil {
			return
		}
	}
}
