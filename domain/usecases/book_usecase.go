package usecases

import (
	"log"

	"github.com/giofcosta/webapi-with-go/domain/entities"
	"github.com/giofcosta/webapi-with-go/domain/repositories"
)

type BookUseCase interface {
	Get(id int) (*entities.Book, error)
	GetAll() (*[]entities.Book, error)
	Create(book entities.Book) (*entities.Book, error)
	Update(book entities.Book) (*entities.Book, error)
	Delete(id int) (*entities.Book, error)
}

type bookUseCase struct {
	repository repositories.BookRepository
}

func NewBookUseCase(repository repositories.BookRepository) BookUseCase {
	return &bookUseCase{repository: repository}
}

func (uc *bookUseCase) Get(id int) (*entities.Book, error) {
	data, err := uc.repository.GetBook(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *bookUseCase) GetAll() (*[]entities.Book, error) {
	data, err := uc.repository.GetBooks()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return data, nil
}

func (uc *bookUseCase) Create(book entities.Book) (*entities.Book, error) {
	data, err := uc.repository.CreateBook(book)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return data, nil
}

func (uc *bookUseCase) Update(book entities.Book) (*entities.Book, error) {
	data, err := uc.repository.UpdateBook(book)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return data, nil
}

func (uc *bookUseCase) Delete(id int) (*entities.Book, error) {
	data, err := uc.repository.DeleteBook(id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return data, nil
}
