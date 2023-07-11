package usecase

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/repository"
)

type BookUsecase interface {
	Read() ([]entity.Book, error)
	Create(judul, pengarang, penerbit, tahunterbit string, status bool) (*entity.Book,error)
	Update(judul, pengarang, penerbit, tahunterbit string, status bool, id int) (*entity.Book, error)
	Delete(id int) error
}

type bookUsecase struct {
    repo repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) bookUsecase {
	return bookUsecase{
			repo: bookRepository,
		}
}

func (u *bookUsecase) Read() ([]entity.Book, error) {
	
	result, err := u.repo.Read()
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (u *bookUsecase) Create(judul, pengarang, penerbit, tahunterbit string, status bool) (*entity.Book, error) {

	book := &entity.Book{
		Judul: judul,
		Pengarang: pengarang,
		Penerbit: penerbit,
		TahunTerbit: tahunterbit,
		Status: status,
	}

	err := u.repo.Create(book)

	if err != nil {
		return nil, err
	}

	return book, nil

}


func (u *bookUsecase) Update(judul, pengarang, penerbit, tahunterbit string, status bool, id int) (*entity.Book, error) {
	
	book := &entity.Book{
		Judul: judul,
		Pengarang: pengarang,
		Penerbit: penerbit,
		TahunTerbit: tahunterbit,
		Status: status,
	}

	result, err := u.repo.Update(book,id)

	if err != nil {
		return nil, err
	}

	return result, nil

	

}

func (u *bookUsecase) Delete(id int) error {
	err := u.repo.Delete(id)
	return err
}