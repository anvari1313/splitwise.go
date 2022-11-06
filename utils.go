package splitwise

import (
	"fmt"
	"reflect"
)

func NewErrFailConversion(value interface{}) error {
	return fmt.Errorf("can't convert value %v", value)
}

func buildStructValues(structFields []reflect.StructField, values []interface{}) (interface{}, error) {
	typ := reflect.StructOf(structFields)

	instance := reflect.New(typ).Elem()
	for i := 0; i < instance.NumField(); i++ {
		varType := instance.Type().Field(i).Type
		switch varType.Kind() {
		case reflect.Int:
			v, ok := (values[i]).(int)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetInt(int64(v))
		case reflect.Int32:
			v, ok := (values[i]).(int32)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetInt(int64(v))
		case reflect.Int64:
			v, ok := (values[i]).(int64)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetInt(v)
		case reflect.Uint:
			v, ok := (values[i]).(uint)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetUint(uint64(v))
		case reflect.Uint32:
			v, ok := (values[i]).(uint32)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetUint(uint64(v))
		case reflect.Uint64:
			v, ok := (values[i]).(uint64)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetUint(v)
		case reflect.Float32:
			v, ok := (values[i]).(float32)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetFloat(float64(v))
		case reflect.Float64:
			v, ok := (values[i]).(float64)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetFloat(v)
		case reflect.Bool:
			v, ok := (values[i]).(bool)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetBool(v)
		case reflect.String:
			v, ok := (values[i]).(string)
			if !ok {
				return nil, NewErrFailConversion(values[i])
			}
			instance.Field(i).SetString(v)
		default:
			return nil, fmt.Errorf("unexpected type %s", varType.Kind())
		}
	}

	return instance.Addr().Interface(), nil
}

func mergeStructFields(inputs ...interface{}) ([]reflect.StructField, []interface{}) {
	var fields []reflect.StructField
	var finalValues []interface{}

	for _, input := range inputs {
		values := reflect.Indirect(reflect.ValueOf(input))
		types := values.Type()

		for i := 0; i < values.NumField(); i++ {
			val := values.Field(i)
			typ := types.Field(i)
			fields = append(fields, reflect.StructField{
				Name: typ.Name,
				Type: reflect.TypeOf(val.Interface()),
				Tag:  reflect.StructTag(string(typ.Tag)),
			})
			finalValues = append(finalValues, val.Interface())
		}
	}

	return fields, finalValues
}

func MergeStructs(inputs ...interface{}) (interface{}, error) {
	fields, values := mergeStructFields(inputs...)

	return buildStructValues(fields, values)
}
