package main

import (
	"encoding/json"
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader"
)

func main() {
	reader := vzReader.VzReader{}

	err := reader.Init()
	if err != nil {
		fmt.Printf("Failed to init reader, err: %v\n", err)
		return
	}

	startTime := "-30m"
	traceIds := []string{"loadtest31CA66B2809A1328292126BA", "loadtest4551B1D31EE68AB61C2F641B", "loadtest56BC003751C65BB8B7088835"}
	data, err := reader.GetHTTPRawData(traceIds, startTime)
	if err != nil {
		fmt.Printf("Failed to get raw data, err: %v\n", err)
		return
	}

	fmt.Printf("DataStats: %v\n", data.ResultStats)
	fmt.Printf("DataResults: \n---\n")
	for _, result := range data.Results {
		obj, _ := json.Marshal(result)
		fmt.Printf("Trace ID: %s\nTrace Data: %s\n---\n", result.TraceId, obj)
	}
}
