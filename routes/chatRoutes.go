package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/chatongo/backend/controllers"
	websockets "github.com/shuklaritvik06/chatongo/backend/websocket"
)

func ChatRoutes(r *mux.Router) {
	r.HandleFunc("/room/{id}", websockets.RoomRouteHandler)
	r.HandleFunc("/room/create/{id}", websockets.CreateRoom).Methods("POST")
	r.HandleFunc("/room/{id}/chats", controllers.GetChats).Methods("GET")
}
