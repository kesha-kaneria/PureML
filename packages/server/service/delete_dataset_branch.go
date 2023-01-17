package service

import (
	"github.com/PureML-Inc/PureML/server/models"
)

// DeleteDatasetBranch godoc
// @Security ApiKeyAuth
// @Summary Delete a branch of a dataset
// @Description Delete a branch of a dataset
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/delete [delete]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
func DeleteDatasetBranch(request *models.Request) *models.Response {
	return nil
}