package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"px.dev/pixie/src/api/go/pxapi/types"
	"strconv"
)

func GetDataByIdx(tag string, datatypeName string, r *types.Record) interface{} {
	var retVal any = nil
	var strRetVal, _ = GetStringFromRecord(tag, r)
	switch datatypeName {
	case "STRING":
		retVal, _ = GetStringFromRecord(tag, r)
	case "TIME64NS":
		retVal, _ = GetTimestampFromRecord(tag, r)
	case "BOOLEAN":
		retVal, _ = GetBooleanFromRecord(tag, r)
	case "INT64", "UINT128":
		retVal, _ = GetIntegerFromRecord(tag, r)
	case "FLOAT64":
		retVal, _ = GetFloatFromRecord(tag, r)
	case "DATA_TYPE_UNKNOWN":
		retVal, _ = GetStringFromRecord(tag, r)
	}

	jsonRetVal := map[string]interface{}{}

	err := json.Unmarshal([]byte(*strRetVal), &jsonRetVal)
	if err != nil {
		println(err)
		return retVal
	}

	return jsonRetVal

}

func GetStringFromRecord(key string, r *types.Record) (*string, error) {
	v := r.GetDatum(key)
	if v == nil {
		return nil, fmt.Errorf(fmt.Sprintf("key %s not found", key))
	}
	return ToPtr[string](v.String()), nil
}

func GetFloatFromRecord(key string, r *types.Record) (*float64, error) {
	dCasted := r.GetDatum(key).(*types.Float64Value)
	floatVal := Round(dCasted.Value(), 12)
	return &floatVal, nil
}

func GetIntegerFromRecord(key string, r *types.Record) (*int, error) {
	s, e := GetStringFromRecord(key, r)
	if s == nil {
		return nil, e
	}
	i, e := GetIntegerFromString(*s)
	return ToPtr[int](i), nil
}

func GetBooleanFromRecord(key string, r *types.Record) (*bool, error) {
	s, e := GetStringFromRecord(key, r)
	if s == nil || e != nil {
		return nil, e
	}
	boolValue, e := strconv.ParseBool(*s)
	return ToPtr[bool](boolValue), e
}

func GetTimestampFromRecord(key string, r *types.Record) (*string, error) {
	t, e := GetStringFromRecord(key, r)
	if t == nil || e != nil {
		return nil, e
	}
	strValue := string(*t)
	return ToPtr[string](strValue), nil
}

func GetFloatFromString(k string, b int) (float64, error) {
	return strconv.ParseFloat(k, b)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func IsEmpty(v string) bool {
	return len(v) == 0
}

// Round Rounds to nearest like 12.3456 -> 12.35
func Round(val float64, precision int) float64 {
	return math.Round(val*(math.Pow10(precision))) / math.Pow10(precision)
}

func GetIntegerFromString(k string) (int, error) {
	return strconv.Atoi(k)
}

func ToPtr[T any](arg T) *T {
	return &arg
}
