package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/giofcosta/webapi-with-go/api/handlers"
)

func ConfigureBookRoutes(router *gin.Engine, handler handlers.BookHandler) *gin.Engine {

	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", handler.ShowBook)
			books.GET("/", handler.ShowBooks)
			books.POST("/", handler.CreateBook)
			books.PUT("/", handler.UpdateBook)
			books.DELETE("/:id", handler.DeleteBook)
		}
	}

	return router
}
