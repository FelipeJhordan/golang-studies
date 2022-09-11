package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

func ConfigureRoutes(router *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, routesLogin)
	for _, route := range routes {

		if route.NeedAuthentication {
			router.HandleFunc(route.URI, middlewares.Authenticate(middlewares.Authenticate(route.Function))).Methods(route.Method)
		}
		router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return router
}
