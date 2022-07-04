package repositories

import (
	"api/src/models"
	"database/sql"
)

type UsersRepository struct {
	db *sql.DB
}

func CreateNewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (u UsersRepository) Create(user models.User) (uint64, error) {
	return 0, nil
}
