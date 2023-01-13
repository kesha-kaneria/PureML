package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// VerifyDatasetHashStatus godoc
// @Security ApiKeyAuth
// @Summary Verify dataset hash status
// @Description Verify dataset hash status to determine if dataset is already uploaded
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/hash-status [post]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param hash body string true "Hash value"
func VerifyDatasetHashStatus(request *models.Request) *models.Response {
	datasetName := request.GetDatasetName()
	datasetUUID := request.GetDatasetUUID()
	if datasetName == "" {
		return models.NewDataResponse(http.StatusOK, true, "Hash validity (True - does not exist in db)")
	}
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := datastore.GetDatasetAllVersions(datasetUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := true
	for _, version := range versions {
		if version.Hash == hashValue {
			response = false
			break
		}
	}
	return models.NewDataResponse(http.StatusOK, response, "Hash validity (True - does not exist in db)")
}
