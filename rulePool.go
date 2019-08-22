package validator

import (
	"errors"
	"mime/multipart"
	"reflect"
)

//校验函数
type FN func(field string, value interface{}, rule *Rule) (bool, error)

//规则池
var rulePool map[string]FN = map[string]FN{}

func init() {
	rulePool["required"] = func(field string, value interface{}, rule *Rule) (bool, error) {
		if value == nil {
			return false, nil
		}

		if _, ok := value.(multipart.File); ok {
			return true, nil
		}

		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if rv.Len() == 0 {
				return false, nil
			}
		case reflect.Int:
			if isEmpty(value.(int)) {
				return false, nil
			}
		case reflect.Int8:
			if isEmpty(value.(int8)) {
				return false, nil
			}
		case reflect.Int16:
			if isEmpty(value.(int16)) {
				return false, nil
			}
		case reflect.Int32:
			if isEmpty(value.(int32)) {
				return false, nil
			}
		case reflect.Int64:
			if isEmpty(value.(int64)) {
				return false, nil
			}
		case reflect.Float32:
			if isEmpty(value.(float32)) {
				return false, nil
			}
		case reflect.Float64:
			if isEmpty(value.(float64)) {
				return false, nil
			}
		case reflect.Uint:
			if isEmpty(value.(uint)) {
				return false, nil
			}
		case reflect.Uint8:
			if isEmpty(value.(uint8)) {
				return false, nil
			}
		case reflect.Uint16:
			if isEmpty(value.(uint16)) {
				return false, nil
			}
		case reflect.Uint32:
			if isEmpty(value.(uint32)) {
				return false, nil
			}
		case reflect.Uint64:
			if isEmpty(value.(uint64)) {
				return false, nil
			}
		case reflect.Uintptr:
			if isEmpty(value.(uintptr)) {
				return false, nil
			}
		default:
			return false, errors.New("invalid type for required rule")

		}
		return true, nil
	}
}
