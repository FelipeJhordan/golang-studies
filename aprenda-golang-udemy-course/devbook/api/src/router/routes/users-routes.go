package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                "/usuarios",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/usuarios",
		Method:             http.MethodGet,
		Function:           controllers.ListUsers,
		NeedAuthentication: false,
	},
	{
		URI:                "/usuarios/{userId}",
		Method:             http.MethodGet,
		Function:           controllers.FindUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/usuarios/{userId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/usuarios/{userId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: false,
	},
}
