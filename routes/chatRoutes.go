package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/chatongo/backend/controllers"
	"github.com/shuklaritvik06/chatongo/backend/utils"
	websockets "github.com/shuklaritvik06/chatongo/backend/websocket"
)

func ChatRoutes(r *mux.Router) {
	new_room := websockets.NewRoom(utils.GenerateUID().String())
	go new_room.Run()
	r.HandleFunc("/room/{id}", new_room.ServeHTTP)
	r.HandleFunc("/room/{id}/chats", controllers.GetChats).Methods("GET")
}
