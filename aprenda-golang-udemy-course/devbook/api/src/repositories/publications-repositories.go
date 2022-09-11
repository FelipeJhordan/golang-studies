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
