package zk_operator

import (
	"encoding/json"
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/api/models"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/api/utils"
)

type OperatorAPIHandler interface {
	GetClusterContext() (*models.ClusterContext, error)
}

type operatorAPIHandler struct {
}

func New() OperatorAPIHandler {
	return &operatorAPIHandler{}
}

func (c *operatorAPIHandler) GetClusterContext() (*models.ClusterContext, error) {
	url := clusterContextURL
	method := clusterContextMethod

	fmt.Printf("GetClusterContext : [%s] %s\n", method, url)

	body, err := utils.Request(method, url)
	responseModel := models.ResponseModel[models.ClusterContext]{}

	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel.Payload, nil
}
