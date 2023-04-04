package routes

import (
	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/chatongo/backend/controllers"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", controllers.Login)
	r.HandleFunc("/auth/register", controllers.Register)
	r.HandleFunc("/refresh", controllers.Refresh)
}
