package models

import (
	"encoding/json"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/utils"
	"px.dev/pixie/src/api/go/pxapi/types"
	"px.dev/pixie/src/api/proto/vizierpb"
)

func ConvertPixieDataToItemStore[itemType ItemType](r *types.Record) itemType {
	var itemStore itemType
	mapObject := map[string]interface{}{}
	for i := 0; i < len(r.Data); i++ {
		tag := r.TableMetadata.ColInfo[i].Name
		datatypeName := vizierpb.DataType_name[int32(r.TableMetadata.ColInfo[i].Type)]
		value := utils.GetDataByIdx(tag, datatypeName, r)
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
