package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Route{
	{
		URI:                "/publicacoes",
		Method:             http.MethodPost,
		Function:           controllers.CreatePublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publicacoes",
		Method:             http.MethodGet,
		Function:           controllers.FindPublications,
		NeedAuthentication: true,
	},
	{
		URI:                "/publicacoes/{publicationId}",
		Method:             http.MethodGet,
		Function:           controllers.FindPublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publicacoes/{publicationId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdatePublications,
		NeedAuthentication: true,
	},
	{
		URI:                "/publicacoes/{publicationId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeletPublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publicacoes/{publicationId}/curtir",
		Method:             http.MethodPost,
		Function:           controllers.LikePublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publicacoes/{publicationId}/descurtir",
		Method:             http.MethodPost,
		Function:           controllers.UnlikePublication,
		NeedAuthentication: true,
	},
}
