package models

import (
	"time"

	"github.com/giofcosta/webapi-with-go/domain/entities"
	"gorm.io/gorm"
)

type BookModel struct {
	ID          int     `gorm:"primary_key;column:id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_Price"`
	Author      string  `json:"author"`
	ImageURL    string  `json:"img_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromBookEntity(e entities.Book) *BookModel {
	return &BookModel{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		MediumPrice: e.MediumPrice,
		Author:      e.Author,
		ImageURL:    e.ImageURL,
	}
}

func (b *BookModel) ToEntity() *entities.Book {
	return &entities.Book{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
		MediumPrice: b.MediumPrice,
		Author:      b.Author,
		ImageURL:    b.ImageURL,
	}
}
