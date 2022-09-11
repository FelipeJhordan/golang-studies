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
	statement, erro := repository.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
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

func (repository UsersRepository) FindByEmail(email string) (models.User, error) {
	line, erro := repository.db.Query("select id, password from users where email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}

	defer line.Close()

	var user models.User
	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

func (repository UsersRepository) FollowUser(userTargetId, followerId uint64) error {
	statement, erro := repository.db.Prepare(
		"insert ignore into followers (userId, followerId) values (?, ?)",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userTargetId, followerId); erro != nil {
		return erro
	}

	return nil
}

func (repository UsersRepository) UnfollowUser(userTargetId, followerId uint64) error {
	statement, erro :=
		repository.db.Prepare(
			"delete from followers where userId = ? and followerId = ?",
		)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userTargetId, followerId); erro != nil {
		return erro
	}

	return nil
}

func (repository UsersRepository) FindFollowers(userId uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(
		`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers f on u.id = f.followerId where f.userId = ?`, userId,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

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

func (repository UsersRepository) FindFollowing(userId uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers f on u.id = f.userId  where f.followerId = ?
	`, userId)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

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

func (repository UsersRepository) FindPassword(userId uint64) (string, error) {
	line, erro := repository.db.Query("select password from users where id = ?", userId)
	if erro != nil {
		return "", erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}

func (repository UsersRepository) UpdatePassword(userId uint64, password string) error {
	statement, erro := repository.db.Prepare("update users set password = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(password, userId); erro != nil {
		return erro
	}

	return nil
}
