package validator

import (
	"fmt"
	"reflect"
)

// isEmpty check a type is Zero
func isEmpty(value interface{}) bool {
	rt := reflect.TypeOf(value)
	if rt == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return rv.Len() == 0
	}
	return reflect.DeepEqual(value, reflect.Zero(rt).Interface())
}

//转字符串
func toString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		str = fmt.Sprintf("%v", v)
	}
	return str
}
