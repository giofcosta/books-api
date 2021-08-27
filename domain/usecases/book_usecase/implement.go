package usecases

import (
	"log"

	"github.com/giofcosta/webapi-with-go/domain/entities"
	repositories "github.com/giofcosta/webapi-with-go/domain/repositories/book_repository"
)

type bookUseCase struct {
	repository repositories.BookRepository
}

func NewBookUseCase(repository repositories.BookRepository) BookUseCase {
	return &bookUseCase{repository: repository}
}

func (uc *bookUseCase) Get(id int) (*entities.Book, error) {
	data, err := uc.repository.GetBook(id)
	if err != nil {
		log.Fatal(err.Error())
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
