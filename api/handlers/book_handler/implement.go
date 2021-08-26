package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/giofcosta/webapi-with-go/domain/entities"
	usecases "github.com/giofcosta/webapi-with-go/domain/usecases/book_usecase"
)

type bookHandler struct {
	usecase usecases.BookUseCase
}

func NewBookHandler(usecase usecases.BookUseCase) BookHandler {
	return &bookHandler{usecase: usecase}
}

func (h *bookHandler) ShowBook(c *gin.Context) {
	id, err := GetId(c)
	if err != nil {
		return
	}

	book, err := h.usecase.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, entities.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, entities.SuccessResponse{Data: book})
}

func (h *bookHandler) CreateBook(c *gin.Context) {

	book, err := BindJson(c)
	if err != nil {
		return
	}

	newBook, err := h.usecase.Create(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Message: "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, entities.SuccessResponse{Data: newBook})
}

func (h *bookHandler) ShowBooks(c *gin.Context) {
	books, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Message: "cannot list books: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entities.SuccessResponse{Data: books})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {

	book, err := BindJson(c)
	if err != nil {
		return
	}

	updatedBook, err := h.usecase.Update(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Message: "cannot update book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entities.SuccessResponse{Data: updatedBook})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	id, err := GetId(c)
	if err != nil {
		return
	}

	err = h.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Message: "cannot delete the book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entities.SuccessResponse{Data: fmt.Sprintf("id:%d. successfully deleted", id)})
}

func GetId(c *gin.Context) (int, error) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Message: "ID has to be an integer",
		})
	}

	return newid, err
}

func BindJson(c *gin.Context) (entities.Book, error) {
	var book entities.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Message: "cannot bind JSON: " + err.Error(),
		})
	}

	return book, err
}
