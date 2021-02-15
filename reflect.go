package injector

import "reflect"

func isStructPtr(val interface{}) bool {
	if val == nil {
		return false
	}

	targetType := reflect.TypeOf(val)
	if targetType.Kind() != reflect.Ptr {
		return false
	}

	return targetType.Elem().Kind() == reflect.Struct
}
