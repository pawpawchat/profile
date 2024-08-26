package validation

import (
	"reflect"
	"time"
)

func GetZeroFields(v interface{}) []string {
	objValue := reflect.ValueOf(v)

	if objValue.Kind() == reflect.Pointer {
		objValue = objValue.Elem()
	}
	if objValue.Kind() != reflect.Struct {
		return nil
	}

	zeroFields := make([]string, 0)
	objType := objValue.Type()

	for idx := range objValue.NumField() {
		field := objValue.Field(idx)

		if !field.CanInterface() {
			continue
		}

		if field.Kind() == reflect.Interface {
			field = field.Elem()
		}

		if field.Kind() == reflect.Pointer {
			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			time.Sleep(10 * time.Millisecond)

			fields := GetZeroFields(field.Interface())
			if len(fields) != 0 {
				zeroFields = append(zeroFields, fields...)
			}
			continue
		}

		if field.IsZero() {
			zeroFields = append(zeroFields, objType.Field(idx).Name)
		}
	}

	return zeroFields
}
