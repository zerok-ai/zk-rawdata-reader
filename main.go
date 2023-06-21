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
	traceIds := []string{"0993ac14ad4eb454eb0bf7559b646682", "0eba241de441af9c2c8873c1c75e9fb6", "1385e0516ff5b6d6d5ca509bdf5cda57"}
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
