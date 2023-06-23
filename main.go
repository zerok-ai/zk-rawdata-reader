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
	traceIds := []string{"loadtest1DFED08C71A1892B3EC26275", "loadtest1F7ED52CACC342BA8DF01601", "loadtest2AAAC9D8E0438EA42B641C4B"}
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
