package migrations

import (
	"fmt"
	"log"

	"github.com/giofcosta/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(models.Book{})
	if err != nil {
		log.Fatal(fmt.Sprintf("%s with error %s", "Migration failed", db.Error))
		return
	}

	log.Print("Migrations was Successfully!!")
}
