package repository

import (
	"github.com/AvinFajarF/internal/entity"
)

type BookRepository interface{
	Read() ([]entity.Book, error)
	Create(books *entity.Book) error
	Update(books *entity.Book, id int) (*entity.Book, error)
	Delete(id int) error
}


func (repo *PostgresUserRepository) Read() ([]entity.Book, error) {
	var books []entity.Book
	result := repo.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (repo *PostgresUserRepository) Create(books *entity.Book) error {
	return repo.db.Create(books).Error
}

func (repo *PostgresUserRepository) Update(books *entity.Book, id int) (*entity.Book ,error) {
	var book entity.Book

	bookFind := repo.db.First(&book, id).Updates(&books)

	if bookFind.Error != nil {
		return nil, bookFind.Error
	}

	return &book, nil
}

func (repo *PostgresUserRepository) Delete(id int) error {
	var book *entity.Book
	return repo.db.Delete(&book, id).Error
}