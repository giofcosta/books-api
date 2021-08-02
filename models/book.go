package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_Price"`
	Author      string  `json:"author"`
	ImageURL    string  `json:"img_url"`
}
