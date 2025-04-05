package main

import (
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/models"
)

type datatype struct {
	models.HttpRawDataModel
	models.MySQLRawDataModel
	models.PgSQLRawDataModel
	models.UPIDToServiceMapModel
}

func main() {
	//reader := vzReader.VzReader{
	//	CloudAddr:  "px.loadcloud01.getanton.com:443",
	//	ClusterId:  "49b97eb5-42bb-4863-9982-fb5d24808a63",
	//	ClusterKey: "px-api-b5dd5d79-549e-4d65-bf4f-5a7f8840ced1",
	//}
	reader := vzReader.VzReader{
		CloudAddr:  "px.loadcloud01.getanton.com:443",
		ClusterId:  "56d95a4d-47e6-4acb-88cd-2588df9b6176",
		ClusterKey: "px-api-f8d4bc05-381d-469b-ab54-82b7128e58ee",

		//CloudAddr:  "px.devcloud01.getanton.com:443",
		//ClusterId:  "6c213fb3-a773-4d8b-b0ad-d6e1b720da6f",
		//ClusterKey: "px-api-e0593597-de51-44cd-bc72-6cbdb881b2be",
	}
	err := reader.Init()
	if err != nil {
		fmt.Printf("Failed to init reader, err: %v\n", err)
		return
	}

	//startTime := "-1h"
	//traceIds := []string{"1f488b4347162a86fb205b62f1e37acf"}

	//data, _ := reader.GetPgSQLRawData(traceIds, startTime)
	//data, _ := reader.GetMySQLRawData(traceIds, startTime)
	//data, _ := reader.GetHTTPRawData(traceIds, startTime)
	data, _ := reader.GetUPIDToServiceMap()

	fmt.Printf("DataStats: %v\n", data.ResultStats)
	fmt.Printf("DataResults: \n---\n")
	for _, result := range data.Results {
		//obj, _ := json.Marshal(result)
		//fmt.Printf("Trace ID: %s\nTrace Data: %s\n---\n", result.TraceId, obj)
		fmt.Printf("UPID: %s\nService: %s\n---\n", result.UPID, result.Service)
	}
}
