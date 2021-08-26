package usecases

import "github.com/giofcosta/webapi-with-go/domain/entities"

type BookUseCase interface {
	Get(id int) (*entities.Book, error)
	GetAll() ([]*entities.Book, error)
	Create(book entities.Book) (*entities.Book, error)
	Update(book entities.Book) (*entities.Book, error)
	Delete(id int) error
}
