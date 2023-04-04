package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shuklaritvik06/chatongo/backend/config"
)

func main() {
	router := mux.NewRouter()
	config.Configure(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
