package vzReader

import (
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/models"
)

// GetHTTPRawData returns the raw data for HTTP for the given traceIds and startTime
func (config *VzReader) GetHTTPRawData(traceIds []string, startTime string) (*models.TraceResponse[models.HTTP_raw_data], error) {
	// Get the script file path for the protocol
	scriptFilePath, ok := protocolMapping[HTTP]
	if !ok {
		return nil, fmt.Errorf("protocol file missing")
	}

	// Get the script string for the protocol
	scriptStr, err := getScriptStr(scriptFilePath, traceIds, startTime)
	if err != nil {
		fmt.Printf("Failed to resolve template, err: %v\n", err)
		return nil, err
	}

	// Execute the script and get the result
	tm := models.New[models.HTTP_raw_data]()
	result, err := executeScript(config.vzClient, config.ctx, scriptStr, tm)
	if err != nil {
		return nil, err
	}

	// Convert the result to HTTP response
	response := models.VzResponseToTraceResponse[models.HTTP_raw_data](result, tm)
	return response, nil
}
