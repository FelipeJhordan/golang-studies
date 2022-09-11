package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"api/src/security/authentication"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdByToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if userId != userIdByToken {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuario que não seja o seu."))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)

	var password models.Password

	if erro = json.Unmarshal(bodyRequest, &password); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	connection, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer connection.Close()

	repository := repositories.CreateNewUserRepository(connection)

	savedPassword, erro := repository.FindPassword(userId)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(savedPassword, password.Current); erro != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("A senha atual não condiz com a senha salva no banco"))
		return
	}

	hashedPassword, erro := security.Hash(password.New)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePassword(userId, string(hashedPassword)); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
