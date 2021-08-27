package main

import (
	"github.com/giofcosta/webapi-with-go/app"
	"github.com/giofcosta/webapi-with-go/database"

	"github.com/giofcosta/webapi-with-go/server"
)

func main() {
	database.StartDB()
	app := app.NewApplication()
	server := server.NewServer()
	server.Run(app)

}
