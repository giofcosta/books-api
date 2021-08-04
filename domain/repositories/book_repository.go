package repositories

import "github.com/giofcosta/webapi-with-go/domain/entities"

type BookRepository interface {
	GetBook(id int) (*entities.Book, error)
	GetBooks() (*[]entities.Book, error)
	CreateBook(book entities.Book) (*entities.Book, error)
	UpdateBook(book entities.Book) (*entities.Book, error)
	DeleteBook(id int) (*entities.Book, error)
}
