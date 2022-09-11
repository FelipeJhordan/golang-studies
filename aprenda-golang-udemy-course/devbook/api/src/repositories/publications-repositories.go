package repositories

import (
	"api/src/models"
	"database/sql"
)

type PublicationRepository struct {
	db *sql.DB
}

func CreateNewPublicationRepository(db *sql.DB) *PublicationRepository {
	return &PublicationRepository{db}
}

func (repository PublicationRepository) Create(publication models.Publication) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into publications (title, content, authorId) values (?,?,?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(publication.Title, publication.Content, publication.AuthorID)

	if erro != nil {
		return 0, erro
	}

	publicationInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(publicationInsertedId), nil
}

func (repository PublicationRepository) FindById(publicationId uint64) (models.Publication, error) {
	line, erro := repository.db.Query(
		`
			select p.*, u.nick from
			publications p inner join users u
			on u.id = p.authorId where p.id = ?
		`, publicationId,
	)

	if erro != nil {
		return models.Publication{}, erro
	}

	defer line.Close()

	var publication models.Publication

	if line.Next() {
		if erro = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return models.Publication{}, erro
		}
	}

	return publication, nil
}

func (repository PublicationRepository) Find(userId uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(`
		select distinct p.*, u.nick from publications p
		inner join users u on u.id = p.authorId
		inner join followers f on p.authorId = f.userId
		where u.id = ? or f.followerID = ?
		order by 1 desc
	`, userId, userId)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication
		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository PublicationRepository) Update(publicationId uint64, publication models.Publication) error {
	statement, erro := repository.db.Prepare(`
		update publications set title = ?, content = ? where id = ?
	`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publication.Title, publication.Content, publicationId); erro != nil {
		return erro
	}

	return nil
}

func (repository PublicationRepository) Delete(publicationId uint64) error {
	statement, erro := repository.db.Prepare("delete from publications where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicationId); erro != nil {
		return erro
	}

	return nil
}

func (repository PublicationRepository) FindPublicationsByUser(userId uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(
		`
			select p.*, u.nick from publications p 
			join users u on u.id = p.authorId
			where p.authorId = ?
		`, userId,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication
		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository PublicationRepository) LikePublication(publicationId uint64) error {
	statement, erro := repository.db.Prepare(` update publications set likes = likes + 1 where id = ?`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicationId); erro != nil {
		return erro
	}

	return nil
}
