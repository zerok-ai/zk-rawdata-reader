package vzReader

import (
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/models"
)

// GetHTTPRawData returns the raw data for HTTP for the given traceIds and startTime
func (config *VzReader) GetHTTPRawData(traceIds []string, startTime string) (*models.TraceResponse[models.HttpRawDataModel], error) {
	// Get the script file path for the protocol
	scriptFileName, ok := protocolMapping[HTTP]
	if !ok {
		return nil, fmt.Errorf("protocol file missing")
	}

	// Get the script string for the protocol
	scriptStr, err := getScriptStr(scriptFileName, traceIds, startTime)
	if err != nil {
		fmt.Printf("Failed to resolve template, err: %v\n", err)
		return nil, err
	}

	// Execute the script and get the result
	tm := models.New[models.HttpRawDataModel]()
	result, err := executeScript(config.vzClient, config.ctx, scriptStr, tm)
	if err != nil {
		return nil, err
	}

	// Convert the result to HTTP response
	response := models.VzResponseToTraceResponse[models.HttpRawDataModel](result, tm)
	return response, nil
}
