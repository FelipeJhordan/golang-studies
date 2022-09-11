package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security/authentication"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(bodyRequest, &publication); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	publication.AuthorID = userId

	if erro = publication.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	connection, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer connection.Close()

	repository := repositories.CreateNewPublicationRepository(connection)
	publication.ID, erro = repository.Create(publication)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}
func FindPublications(w http.ResponseWriter, r *http.Request) {
	userIdByToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	connection, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	defer connection.Close()

	repository := repositories.CreateNewPublicationRepository(connection)
	publications, erro := repository.Find(userIdByToken)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publications)

}
func FindPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repository := repositories.CreateNewPublicationRepository(db)
	publication, erro := repository.FindById(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, publication)

}
func UpdatePublications(w http.ResponseWriter, r *http.Request) {

	userIdByToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repository := repositories.CreateNewPublicationRepository(db)
	savedPublication, erro := repository.FindById(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if savedPublication.AuthorID != userIdByToken {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não é sua."))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
	}

	var publicacao models.Publication
	if erro = json.Unmarshal(bodyRequest, &publicacao); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.Update(publicationId, publicacao); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletPublication(w http.ResponseWriter, r *http.Request) {

}
