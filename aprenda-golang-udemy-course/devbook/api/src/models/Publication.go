package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (publication *Publication) Prepare() error {
	if erro := publication.validate(); erro != nil {
		return erro
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if publication.Content == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publication *Publication) format() error {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
	return nil
}
