package app

import (
	handlers "github.com/giofcosta/webapi-with-go/api/handlers/book_handler"
	repositories "github.com/giofcosta/webapi-with-go/data/repositories/book_repository"
	"github.com/giofcosta/webapi-with-go/database"

	usecases "github.com/giofcosta/webapi-with-go/domain/usecases/book_usecase"
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
	repository := repositories.NewBookRepository(database.GetDatabase())
	useCase := usecases.NewBookUseCase(repository)
	handler := handlers.NewBookHandler(useCase)

	return &handler
}
