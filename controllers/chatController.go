package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/chatongo/backend/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetChats(w http.ResponseWriter, r *http.Request) {
	cur, _ := database.GetDB().Database("chats").Collection(mux.Vars(r)["id"]).Find(context.Background(), bson.D{{}})
	var chats []bson.M
	if chats == nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No Documents",
		})
		return
	}
	cur.All(context.Background(), &chats)
	json.NewEncoder(w).Encode(chats)
}
