package router

import (
	"test/db"
	"test/handler"

	"github.com/gorilla/mux"
)

type Router struct {
	R  *mux.Router
	DB db.Database
}

func (r *Router) Setup() {
	r.R = mux.NewRouter()
	r.DB, _ = db.NewDB()
	h := handler.NewHTTPHandler(r.DB)
	r.R.HandleFunc("/user/create", h.CreateNewUser).Methods("POST")
	r.R.HandleFunc("/user/{id}", h.GetUserProfile).Methods("GET")
}
func NewRouter() Router {
	var r Router
	r.Setup()
	return r
}
