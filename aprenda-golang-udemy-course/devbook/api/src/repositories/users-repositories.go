package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository UsersRepository) ListUsers(nameOrNick string) ([]models.User, error) {
	nameOrNickLikeStatement := fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name like ? or nick like ?",
		nameOrNickLikeStatement, nameOrNickLikeStatement,
	)

	if erro != nil {
		return nil, erro
	}

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository UsersRepository) FindUser(userId uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?;", userId,
	)

	if erro != nil {
		return models.User{}, erro
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository UsersRepository) Update(id uint64, user models.User) error {
	statatment, erro :=
		repository.db.Prepare(
			"update users set name = ?, nick = ?, email = ? where id = ? ",
		)

	if erro != nil {
		return erro
	}

	defer statatment.Close()

	if _, erro = statatment.Exec(user.Name, user.Nick, user.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (repository UsersRepository) Deletar(userId uint64) error {
	statatment, erro := repository.db.Prepare("delete from users where id = ?")

	if erro != nil {
		return erro
	}

	defer statatment.Close()

	if _, erro = statatment.Exec(userId); erro != nil {
		return erro
	}

	return nil

}