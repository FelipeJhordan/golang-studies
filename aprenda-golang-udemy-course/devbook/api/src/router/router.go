package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func GenerateRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigureRoutes(router)
}
