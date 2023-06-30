package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface{
	Save(user *entity.User) error
	Login(email, password string, user *entity.User) error}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
        db: db,
    }
}

func (repo *PostgresUserRepository) Save(user *entity.User) error {
	user.ID = uuid.New().String()
	return repo.db.Create(user).Error
}

func (repo *PostgresUserRepository) Login(email, password string, user *entity.User) error {
	return repo.db.Where("email = ?", email).First(user).Error
}