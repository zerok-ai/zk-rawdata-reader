package models

import (
	"encoding/json"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/utils"
	"px.dev/pixie/src/api/go/pxapi/types"
	"px.dev/pixie/src/api/proto/vizierpb"
)

func ConvertPixieDataToItemStore[itemType ItemType](records *types.Record) itemType {
	var itemStore itemType
	mapObject := map[string]interface{}{}
	recordCount := len(records.Data)
	for i := 0; i < recordCount; i++ {
		tag := records.TableMetadata.ColInfo[i].Name
		datatypeName := vizierpb.DataType_name[int32(records.TableMetadata.ColInfo[i].Type)]
		value := utils.GetDataByIdx(tag, datatypeName, records)
		mapObject[tag] = value
	}

	if mapObject["time_"] != nil {
		mapObject["time"] = mapObject["time_"]
		delete(mapObject, "time_")
	}

	jsonStr, err := json.Marshal(mapObject)
	if err != nil {
		println(err)
	}

	err = json.Unmarshal(jsonStr, &itemStore)
	if err != nil {
		println(err)
	}

	return itemStore
}
