package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Env struct {
	Db *gorm.DB
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	return r
}

func SetRoutes(router *mux.Router, env *Env) {
	router.HandleFunc("/parents", env.GetParent).Methods("GET")
	router.HandleFunc("/parent", env.PostParent).Methods("POST")

	router.HandleFunc("/websocket/echo", env.Echo)
	router.HandleFunc("/websocket/children", env.RetrieveChildren)
}
