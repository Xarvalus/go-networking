package core

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}
}
