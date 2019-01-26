package core

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const ServerPort = ":8000"

func StartServer (router *mux.Router) {
	log.Fatal(http.ListenAndServe(ServerPort, router))
}
