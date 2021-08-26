package handlers

import (
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	ShowBook(c *gin.Context)
	CreateBook(c *gin.Context)
	ShowBooks(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}
