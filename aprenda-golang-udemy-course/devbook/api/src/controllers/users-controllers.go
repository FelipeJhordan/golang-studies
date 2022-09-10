package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if erro = user.Prepare("register"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
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
	nameOrNickQuery := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.CreateNewUserRepository(db)

	users, erro := repository.ListUsers(nameOrNickQuery)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}
func FindUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	userId, erro := strconv.ParseUint(pathParams["userId"], 10, 64)

	if erro != nil {
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

	user, erro := repository.FindUser(userId)
	if erro != nil {

		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if user.ID == 0 {
		responses.JSON(w, http.StatusNotFound, nil)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("edit"); erro != nil {
		responses.Error(w, http.StatusBadGateway, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.CreateNewUserRepository(db)
	if erro = repository.Update(userId, user); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	if erro := repository.Deletar(userId); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
