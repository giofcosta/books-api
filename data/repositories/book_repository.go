package repositories

import (
	"errors"

	common "github.com/giofcosta/webapi-with-go/common/errors"
	"github.com/giofcosta/webapi-with-go/data/models"
	"github.com/giofcosta/webapi-with-go/domain/entities"
	"github.com/giofcosta/webapi-with-go/domain/repositories"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repositories.BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) GetBook(id int) (*entities.Book, error) {
	var book models.BookModel

	err := r.db.First(&book, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrBookNotFound
		}

		return nil, err
	}

	entity := book.ToEntity()
	return entity, err
}

func (r *bookRepository) GetBooks() (*[]entities.Book, error) {
	var books []models.BookModel
	var entities []entities.Book
	err := r.db.Find(&books).Error

	if err != nil {
		return nil, err
	}

	for _, v := range books {
		entities = append(entities, *v.ToEntity())
	}

	return &entities, err
}

func (r *bookRepository) CreateBook(book entities.Book) (*entities.Book, error) {
	model := models.FromBookEntity(book)
	err := r.db.Create(model).Error

	if err != nil {
		return nil, err
	}

	return model.ToEntity(), err
}

func (r *bookRepository) UpdateBook(book entities.Book) (*entities.Book, error) {
	model := models.FromBookEntity(book)
	err := r.db.Save(&model).Error

	if err != nil {
		return nil, err
	}
	//TODO: Automapper
	//var ebook entities.Book
	//automapper.MapLoose(model, &ebook)

	return model.ToEntity(), err
}

func (r *bookRepository) DeleteBook(id int) (*entities.Book, error) {
	var model models.BookModel
	err := r.db.Delete(&model, id).Error

	if err != nil {
		return nil, err
	}

	return model.ToEntity(), err
}
