package repositories

import "github.com/giofcosta/webapi-with-go/domain/entities"

type bookRepositoryMock struct {
	MockGetBook    func(id int) (*entities.Book, error)
	MockGetBooks   func() (*[]entities.Book, error)
	MockCreateBook func(book entities.Book) (*entities.Book, error)
	MockUpdateBook func(book entities.Book) (*entities.Book, error)
	MockDeleteBook func(id int) (*entities.Book, error)
}

func NewBookRepositoryMock() *bookRepositoryMock {
	return &bookRepositoryMock{}
}

func (fake *bookRepositoryMock) GetBook(id int) (*entities.Book, error) {
	return fake.MockGetBook(id)
}

func (fake *bookRepositoryMock) GetBooks() (*[]entities.Book, error) {
	return fake.MockGetBooks()
}

func (fake *bookRepositoryMock) CreateBook(book entities.Book) (*entities.Book, error) {
	return fake.MockCreateBook(book)
}

func (fake *bookRepositoryMock) UpdateBook(book entities.Book) (*entities.Book, error) {
	return fake.MockUpdateBook(book)
}

func (fake *bookRepositoryMock) DeleteBook(id int) (*entities.Book, error) {
	return fake.MockDeleteBook(id)
}
