package vzReader

import (
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/models"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/pxl"
)

const (
	UPIDToServiceMapPxl = "getUPIDToServiceMap.pxl"
)

func (config *VzReader) GetUPIDToServiceMap() (*models.TraceResponse[models.UPIDToServiceMapModel], error) {
	scriptStr, err := pxl.GetPxlFileContent(UPIDToServiceMapPxl)
	if err != nil {
		fmt.Printf("failed to get script file content %s\n", UPIDToServiceMapPxl)
		return nil, err
	}

	// Execute the script and get the result
	tm := models.New[models.UPIDToServiceMapModel]()
	result, err := executeScript(config.vzClient, config.ctx, scriptStr, tm)
	if err != nil {
		return nil, err
	}

	// Convert the result to HTTP response
	response := models.VzResponseToTraceResponse[models.UPIDToServiceMapModel](result, tm)
	return response, nil
}
