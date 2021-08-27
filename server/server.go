package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/giofcosta/webapi-with-go/app"
	"github.com/giofcosta/webapi-with-go/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run(app *app.Application) {
	router := routes.ConfigureBookRoutes(s.server, *app.BookHandler)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
