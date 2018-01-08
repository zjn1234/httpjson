package httpjson

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
 * slice only support []int, []bool, []string, []float
 */

var (
	//default space
	sliceDelim rune = ' '
)

func SetDelim(c rune) {
	sliceDelim = c
}

func delimFunc(c rune) bool {
	return sliceDelim == c
}

func setSlice(field reflect.Value, formValue string) error {
	elemType := field.Type().Elem().Kind()
	slice, err := fieldFuncByKind(elemType, formValue, delimFunc)
	if err != nil {
		return err
	}

	field.Set(reflect.ValueOf(slice))
	return nil
}

func fieldFuncByKind(k reflect.Kind, s string, f func(rune) bool) (interface{}, error) {
	stringSlice := strings.FieldsFunc(s, f)
	switch k {
	case reflect.String:
		return stringSlice, nil

	case reflect.Int:
		intSlice := make([]int, len(stringSlice), len(stringSlice))
		for key, value := range stringSlice {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			intSlice[key] = num
		}
		return intSlice, nil

	case reflect.Float32:
		float32Slice := make([]float32, len(stringSlice), len(stringSlice))
		for key, value := range stringSlice {
			num, err := atoFloat32(value)
			if err != nil {
				return nil, err
			}
			float32Slice[key] = num
		}
		return float32Slice, nil

	case reflect.Float64:
		float64Slice := make([]float64, len(stringSlice), len(stringSlice))
		for key, value := range stringSlice {
			num, err := atoFloat64(value)
			if err != nil {
				return nil, err
			}
			float64Slice[key] = num
		}
		return float64Slice, nil

	case reflect.Bool:
		boolSlice := make([]bool, len(stringSlice), len(stringSlice))
		for key, value := range stringSlice {
			num, err := atoBool(value)
			if err != nil {
				return nil, err
			}
			boolSlice[key] = num
		}
		return boolSlice, nil

	default:
		return nil, fmt.Errorf("not support type: %v", k)
	}

}

func atoFloat32(s string) (float32, error) {
	value64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(value64), nil
}
func atoFloat64(s string) (float64, error) {
	value64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return value64, nil
}
func atoBool(s string) (bool, error) {
	valueBool, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}

	return valueBool, nil
}
