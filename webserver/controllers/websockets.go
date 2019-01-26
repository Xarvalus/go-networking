package controllers

import (
	"encoding/json"
	"github.com/Xarvalus/go-networking/webserver/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func webSocketConnect(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// TODO: should not be blindly omitted on production
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("[CONNECTION ERROR]:", err)
		return nil
	}

	return conn
}

func (env *Env) Echo(w http.ResponseWriter, r *http.Request) {
	conn := webSocketConnect(w, r)
	if conn == nil {
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("[READ ERROR]:", err)
			break
		}

		log.Printf("[RECEIVED]: (%d) %s", mt, message)

		// Echoing message back to the client
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("[WRITE ERROR]:", err)
			break
		}
	}
}

func (env *Env) RetrieveChildren(w http.ResponseWriter, r *http.Request) {
	conn := webSocketConnect(w, r)
	if conn == nil {
		return
	}
	defer conn.Close()

	children := env.Db.Find(&[]models.Children{})
	jsonChildren, jsonErr := json.Marshal(children.Value)
	if jsonErr != nil {
		log.Println(jsonErr)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, jsonChildren)
	if err != nil {
		log.Println("[WRITE ERROR]:", err)
	}
}
