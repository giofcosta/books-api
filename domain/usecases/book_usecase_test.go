package usecases_test

import (
	"reflect"
	"testing"

	"github.com/giofcosta/webapi-with-go/domain/entities"
	"github.com/giofcosta/webapi-with-go/domain/repositories"
	"github.com/giofcosta/webapi-with-go/domain/usecases"
)

func TestGet(t *testing.T) {
	r := repositories.NewBookRepositoryMock()
	uc := usecases.NewBookUseCase(r)

	tests := []struct {
		name    string
		id      int
		mock    func()
		want    *entities.Book
		wantErr bool
	}{
		{
			name: "Should be ok",
			id:   1,
			mock: func() {
				r.MockGetBook = func(id int) (*entities.Book, error) {
					return &entities.Book{}, nil
				}
			},
			want:    &entities.Book{},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			book, err := uc.Get(test.id)

			if (err != nil) != test.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, test.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(book, test.want) {
				t.Errorf("Get() = %v, want %v", book, test.want)
			}
		})
	}

}
