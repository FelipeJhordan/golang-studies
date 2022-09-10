package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repository := repositories.CreateNewUserRepository(db)
	user.ID, erro = repository.Create(user)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar usuários"))
}
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuário"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuário"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Usuário"))
}
