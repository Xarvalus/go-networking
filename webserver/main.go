package main

import (
	"github.com/Xarvalus/go-networking/webserver/controllers"
	"github.com/Xarvalus/go-networking/webserver/core"
)

// Simple WebServer
// JSON REST API & Websockets
// With data layer as Postgres & GORM ORM
func main() {
	db := core.Connect()
	core.AutoMigrate(db)
	defer core.Close(db)

	env := &controllers.Env{Db: db}
	r := controllers.InitRouter()
	controllers.SetRoutes(r, env)

	core.StartServer(r)
}
