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
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}",
		Method:             http.MethodGet,
		Function:           controllers.FindUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}/seguir",
		Method:             http.MethodPost,
		Function:           controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}/parar-de-seguir",
		Method:             http.MethodPost,
		Function:           controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}/seguidores",
		Method:             http.MethodGet,
		Function:           controllers.FindFollowers,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}/seguindo",
		Method:             http.MethodGet,
		Function:           controllers.FindFollowing,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}/atualizar-senha",
		Method:             http.MethodPost,
		Function:           controllers.UpdatePassword,
		NeedAuthentication: true,
	},
	{
		URI:                "/usuarios/{userId}/publicacoes",
		Method:             http.MethodGet,
		Function:           controllers.FindPublicationsByUser,
		NeedAuthentication: true,
	},
}
