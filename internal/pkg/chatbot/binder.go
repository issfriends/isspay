package chatbot

import (
	"reflect"
	"strconv"
	"strings"
)

func bind(data map[string]string, obj interface{}, tag string) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {

	}
	val = val.Elem()
	tye := val.Type()

	for i := 0; i < val.NumField(); i++ {
		tag := tye.Field(i).Tag.Get(tag)
		fieldVal, ok := data[tag]
		if ok {
			field := val.Field(i)
			if field.Kind() == reflect.Ptr {
				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}
				field = field.Elem()
			}
			if field.CanSet() {
				setValue(fieldVal, field)
			}
		}
	}
	return nil
}

func setValue(data string, field reflect.Value) error {
	kind := field.Kind()

	if kind >= reflect.Int && kind <= reflect.Int64 {
		val, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(val)
	} else if kind >= reflect.Uint && kind <= reflect.Uint64 {
		val, err := strconv.ParseUint(data, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(val)
	} else if kind == reflect.String {
		field.SetString(data)
	} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
		val, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return err
		}
		field.SetFloat(val)
	} else if kind == reflect.Bool {
		val := false
		data = strings.ToLower(data)
		if data == "yes" || data == "1" || data == "y" || data == "true" {
			val = true
		}
		field.SetBool(val)
	} else if kind == reflect.Slice {
		var newField reflect.Value
		values := strings.Split(data, ",")

		if field.Cap() == 0 {
			newField = reflect.MakeSlice(field.Type(), len(values), len(values))
		}

		for i := 0; i < len(values); i++ {
			if err := setValue(values[i], newField.Index(i)); err != nil {
				return err
			}
		}
		field.Set(newField)
	}

	return nil
}
