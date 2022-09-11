package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"api/src/security/authentication"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
	}

	var user models.User

	if erro := json.Unmarshal(body, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.CreateNewUserRepository(db)

	savedUserDB, erro := repository.FindByEmail(user.Email)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(savedUserDB.Password, user.Password); erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CreateToken(savedUserDB.ID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	fmt.Println(token)
	w.Write([]byte(token))
}
