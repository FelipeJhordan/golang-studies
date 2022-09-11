package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Route{
	{
		URI:                "/publications",
		Method:             http.MethodPost,
		Function:           controllers.CreatePublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications",
		Method:             http.MethodGet,
		Function:           controllers.FindPublications,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Method:             http.MethodGet,
		Function:           controllers.FindPublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdatePublications,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeletPublication,
		NeedAuthentication: true,
	},
}
