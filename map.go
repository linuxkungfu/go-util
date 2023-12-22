package util

import (
	"encoding/json"
	"reflect"
	"strconv"

	logger "github.com/sirupsen/logrus"
)

func MapArrayToInterface(dataMap interface{}, dstInterface interface{}) interface{} {
	// err := mapstructure.Decode(dataMap, dstInterface)
	str, err := json.Marshal(dataMap)
	if err != nil {
		return nil
	} else {
		err = json.Unmarshal(str, dstInterface)
		if err != nil {
			return nil
		}
		return dstInterface
	}
}
func GetIntValueFromMap(key string, dataMap map[string]interface{}) (int, bool) {
	valueInterface, valueOk := dataMap[key]
	if !valueOk {
		return 0, valueOk
	}
	return GetIntValueFromInterface(valueInterface)
}

func GetUint64ValueFromMap(key string, dataMap map[string]interface{}) (uint64, bool) {
	valueInterface, valueOk := dataMap[key]
	if !valueOk {
		return 0, valueOk
	}
	return GetUint64ValueFromInterface(valueInterface)
}

func GetInt64ValueFromMap(key string, dataMap map[string]interface{}) (int64, bool) {
	valueInterface, valueOk := dataMap[key]
	if !valueOk {
		return 0, valueOk
	}
	return GetInt64ValueFromInterface(valueInterface)
}

func GetStringValueFromMap(key string, dataMap map[string]interface{}) (string, bool) {
	valueInterface, valueOk := dataMap[key]
	if !valueOk {
		return "", valueOk
	}
	switch valueInterface.(type) {
	case string:
		return valueInterface.(string), true
	default:
		return "", false
	}
}

func GetIntValueFromInterface(value interface{}) (int, bool) {
	switch value.(type) {
	case int8:
		return int(value.(int8)), true
	case int16:
		return int(value.(int16)), true
	case int32:
		return int(value.(int32)), true
	case int64:
		return int(value.(int64)), true
	case uint8:
		return int(value.(uint8)), true
	case uint16:
		return int(value.(uint16)), true
	case uint32:
		return int(value.(uint32)), true
	case uint64:
		return int(value.(uint64)), true
	case float32:
		return int(value.(float32)), true
	case float64:
		return int(value.(float64)), true
	case string:
		i, err := strconv.Atoi(value.(string))
		if err != nil {
			logger.Warnf("[map][]GetIntValueFromInterface string to int failed:%s", value.(string))
			return 0, false
		} else {
			return i, true
		}
	default:
		logger.Warnf("[map][]GetIntValueFromInterface illegal type:%v", reflect.TypeOf(value).Kind())
		return 0, false
	}
}

func GetUintValueFromInterface(value interface{}) (uint, bool) {
	switch value.(type) {
	case int8:
		return uint(value.(int8)), true
	case int16:
		return uint(value.(int16)), true
	case int32:
		return uint(value.(int32)), true
	case int64:
		return uint(value.(int64)), true
	case uint8:
		return uint(value.(uint8)), true
	case uint16:
		return uint(value.(uint16)), true
	case uint32:
		return uint(value.(uint32)), true
	case uint64:
		return uint(value.(uint64)), true
	case float32:
		return uint(value.(float32)), true
	case float64:
		return uint(value.(float64)), true
	case string:
		ui64, err := strconv.ParseUint(value.(string), 0, 64)
		if err != nil {
			logger.Warnf("[map][]GetUintValueFromInterface failed:%s", err.Error())
			return 0, false
		} else {
			return uint(ui64), true
		}
	default:
		logger.Warnf("[map][]GetUintValueFromInterface illegal type:%v", reflect.TypeOf(value).Kind())
		return 0, false
	}
}

func GetUint64ValueFromInterface(value interface{}) (uint64, bool) {
	switch value.(type) {
	case int8:
		return uint64(value.(int8)), true
	case int16:
		return uint64(value.(int16)), true
	case int32:
		return uint64(value.(int32)), true
	case int64:
		return uint64(value.(int64)), true
	case uint8:
		return uint64(value.(uint8)), true
	case uint16:
		return uint64(value.(uint16)), true
	case uint32:
		return uint64(value.(uint32)), true
	case uint64:
		return uint64(value.(uint64)), true
	case float32:
		return uint64(value.(float32)), true
	case float64:
		return uint64(value.(float64)), true
	case string:
		ui64, err := strconv.ParseUint(value.(string), 0, 64)
		if err != nil {
			logger.Warnf("[map][]GetUint64ValueFromInterface failed:%s", err.Error())
			return 0, false
		} else {
			return ui64, true
		}
	default:
		logger.Warnf("[map][]GetUint64ValueFromInterface illegal type:%v", reflect.TypeOf(value).Kind())
		return 0, false
	}
}

func GetInt64ValueFromInterface(value interface{}) (int64, bool) {
	switch value.(type) {
	case int8:
		return int64(value.(int8)), true
	case int16:
		return int64(value.(int16)), true
	case int32:
		return int64(value.(int32)), true
	case int64:
		return int64(value.(int64)), true
	case uint8:
		return int64(value.(uint8)), true
	case uint16:
		return int64(value.(uint16)), true
	case uint32:
		return int64(value.(uint32)), true
	case uint64:
		return int64(value.(uint64)), true
	case float32:
		return int64(value.(float32)), true
	case float64:
		return int64(value.(float64)), true
	case string:
		i64, err := strconv.ParseInt(value.(string), 0, 64)
		if err != nil {
			logger.Warnf("[map][]GetInt64ValueFromInterface failed:%s", err.Error())
			return 0, false
		} else {
			return i64, true
		}
	default:
		logger.Warnf("[map][]GetInt64ValueFromInterface illegal type:%v", reflect.TypeOf(value).Kind())
		return 0, false
	}
}

func replaceAndMergeInterface(first reflect.Value, second reflect.Value) reflect.Value {
	if first.IsZero() {
		return second
	}
	if first.IsValid() && (first.Kind() != reflect.Struct && first.Kind() != reflect.Slice) {
		return first
	}
	if second.Kind() == reflect.Struct {
		for index := 0; index < second.NumField(); index++ {
			secondSubValue := second.Field(index)
			// logger.Infof("=====%s", second.Type().Field(index).Name)
			firstSubValue := first.FieldByName(second.Type().Field(index).Name)
			if !firstSubValue.IsZero() {
				// logger.Infof("=====%s", firstSubValue.Kind().String())
				if firstSubValue.Kind() == reflect.Struct || firstSubValue.Kind() == reflect.Slice {
					secondSubValue.Set(replaceAndMergeInterface(firstSubValue, secondSubValue))
				} else {
					secondSubValue.Set(firstSubValue)
				}
			}
		}
	} else if second.Kind() == reflect.Slice {
		for index := 0; index < first.Len(); index++ {
			if second.Len() <= index {
				second = reflect.Append(second, first.Index(index))
			} else {
				secondValue := second.Index(index)
				firstValue := first.Index(index)
				if !firstValue.IsZero() {
					if firstValue.Kind() == reflect.Struct || firstValue.Kind() == reflect.Slice {
						secondValue.Set(replaceAndMergeInterface(firstValue, secondValue))
					} else {
						secondValue.Set(firstValue)
					}
				}
			}
		}
	}
	return second
}

// mergeConfig merge custom config and default config
func MergeInterface[T interface{}](first, second *T) *T {
	if reflect.ValueOf(second).IsNil() {
		return first
	} else if reflect.ValueOf(first).IsNil() {
		return second
	}
	result := new(T)
	secondReflectValue := reflect.ValueOf(second).Elem()
	secondReflectType := secondReflectValue.Type()
	firstReflectValue := reflect.ValueOf(first).Elem()
	firstReflectType := firstReflectValue.Type()
	resultRefelectValue := reflect.ValueOf(result).Elem()
	for index := 0; index < secondReflectValue.NumField(); index++ {
		fieldObject := secondReflectValue.Field(index)
		resultFieldValue := resultRefelectValue.FieldByName(secondReflectType.Field(index).Name)
		if resultFieldValue.IsValid() && resultFieldValue.CanSet() {
			resultFieldValue.Set(fieldObject)
		}

	}
	for index := 0; index < firstReflectValue.NumField(); index++ {
		fieldObject := firstReflectValue.Field(index)
		if fieldObject.Kind() == reflect.Invalid {
			continue
		}
		// logger.Infof("=====%s", envReflectType.Field(index).Name)
		resultFieldValue := resultRefelectValue.FieldByName(firstReflectType.Field(index).Name)
		resultFieldValue.Set(replaceAndMergeInterface(fieldObject, resultFieldValue))
	}
	return result
}
