package app

import (
	"github.com/giofcosta/webapi-with-go/api/handlers"
	"github.com/giofcosta/webapi-with-go/data/database"
	"github.com/giofcosta/webapi-with-go/data/repositories"
	"github.com/giofcosta/webapi-with-go/domain/usecases"
)

type Application struct {
	BookHandler *handlers.BookHandler
}

func NewApplication() *Application {
	return &Application{
		BookHandler: buildBooksHandler(),
	}
}

func buildBooksHandler() *handlers.BookHandler {
	db := database.GetDatabase()
	repository := repositories.NewBookRepository(db)
	useCase := usecases.NewBookUseCase(repository)
	handler := handlers.NewBookHandler(useCase)

	return &handler
}
