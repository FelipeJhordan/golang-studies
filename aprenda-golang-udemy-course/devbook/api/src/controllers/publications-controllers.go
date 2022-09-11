package controllers

import (
	db "api/src/bd"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security/authentication"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

}
func FindPublication(w http.ResponseWriter, r *http.Request) {

}
func UpdatePublications(w http.ResponseWriter, r *http.Request) {

}
func DeletPublication(w http.ResponseWriter, r *http.Request) {

}
