package usecase

import "github.com/AvinFajarF/internal/entity"

type BookUsecase interface {
	Read() ([]entity.Book, error)
	Create(books *entity.Book) error
	Update(books *entity.Book, id int) (*entity.Book, error)
	Delete(id int) error
}