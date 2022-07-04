package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.CreateNewUserRepository(db)
	repository.Create(user)
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
