package main

import (
	"encoding/json"
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader"
)

func main() {
	reader := vzReader.VzReader{
		CloudAddr:   "px.avinpx07.getanton.com:443",
		DirectVzId:  "94711f31-f693-46be-91c3-832c0f64b12f",
		DirectVzKey: "px-api-ce1bbae5-49c7-4d81-99e2-0d11865bb5df",
	}
	
	err := reader.Init()
	if err != nil {
		fmt.Printf("Failed to init reader, err: %v\n", err)
		return
	}

	startTime := "-30m"
	traceIds := []string{"03e2a65ac1654e9d91a4f288d2ccadee", "193a67aa0fb2d76ce9d1024539c3cfba", "229e637ca64c0cff0ae2f5effc0fdf40"}
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
