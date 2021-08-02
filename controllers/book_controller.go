package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/giofcosta/webapi-with-go/database"
	"github.com/giofcosta/webapi-with-go/models"
)

func ShowBook(c *gin.Context) {
	id, err := GetId(c)
	if err != nil {
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find the book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	book, err := BindJson(c)
	if err != nil {
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	book, err := BindJson(c)
	if err != nil {
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id, err := GetId(c)
	if err != nil {
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.Delete(&book, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete the book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}

func GetId(c *gin.Context) (*int, error) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
	}

	return &newid, err
}

func BindJson(c *gin.Context) (*models.Book, error) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
	}

	return &book, err
}
