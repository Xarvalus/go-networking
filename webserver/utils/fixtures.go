package main

import (
	"github.com/Xarvalus/go-networking/webserver/core"
	"github.com/Xarvalus/go-networking/webserver/models"
)

func main() {
	db := core.Connect()
	core.AutoMigrate(db)
	defer core.Close(db)


	db.Create(&models.Parent{
		Number: 10,
		Text: "some text",
		Real: 5.75,
		Enum: models.Two,
		Children: []models.Children{
			{Boolean: true},
			{Boolean: false},
		},
	})
}
