package config

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/shuklaritvik06/chatongo/backend/database"
	"github.com/shuklaritvik06/chatongo/backend/routes"
)

func Configure(r *mux.Router) {
	godotenv.Load(".env")
	routes.ChatRoutes(r)
	routes.AuthRoutes(r)
	database.DBConnect()
}
