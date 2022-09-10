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

func (repository UsersRepository) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, nil
	}

	lastInsertedId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedId), nil
}
