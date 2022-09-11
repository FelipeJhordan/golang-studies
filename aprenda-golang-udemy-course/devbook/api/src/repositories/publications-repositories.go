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
