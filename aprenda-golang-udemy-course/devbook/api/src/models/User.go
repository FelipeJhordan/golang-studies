package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	user.format()

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco.")
	}
	if user.Nick == "" {
		return errors.New("O apelido é obrigatório e não pode estar em branco.")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco.")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}

	if step == "register" && user.Password == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco.")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
