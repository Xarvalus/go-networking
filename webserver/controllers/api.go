package controllers

import (
	"encoding/json"
	"github.com/Xarvalus/go-networking/webserver/core"
	"github.com/Xarvalus/go-networking/webserver/models"
	"log"
	"net/http"
)

func (env *Env) GetParent(w http.ResponseWriter, r *http.Request) {
	parents := env.Db.Find(&[]models.Parent{})

	core.JSON(w, parents.Value)
}

func (env *Env) PostParent(w http.ResponseWriter, r *http.Request) {
	/*
	// Example JSON POST request
	{
		"Number": 10,
		"Text": "some text",
		"Real": 5.75,
		"Enum": "THREE"
	}
	 */

	decoder := json.NewDecoder(r.Body)

	var res struct {
		Number uint
		Text   string
		Real   float64
		Enum   models.Enum
	}

	err := decoder.Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	env.Db.Create(&models.Parent{
		Number: res.Number,
		Text: res.Text,
		Real: res.Real,
		Enum: res.Enum,
	})

	core.JSON(w, "ok")
}
