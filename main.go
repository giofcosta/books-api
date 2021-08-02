package main

import (
	"github.com/giofcosta/webapi-with-go/database"
	"github.com/giofcosta/webapi-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
