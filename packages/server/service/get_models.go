package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

func GetAllModels(request *models.Request) *models.Response {
	orgId := request.GetPathParam("orgId")
	allModels, err := datastore.GetAllModels(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allModels, "Models successfully retrieved")
}