package httpjson

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	nilError = errors.New("argument nil error")
)

func Unmarshal(r *http.Request, v interface{}) error {
	if r == nil || v == nil {
		return nilError
	}

	valSet := reflect.ValueOf(v).Elem()
	valType := reflect.TypeOf(v).Elem()
	numField := valType.NumField()

	for i := 0; i < numField; i++ {
		fieldType := valType.Field(i)
		fieldValue := valSet.Field(i)

		key := findMatchKey(fieldType)
		formValue := r.FormValue(key)
		if formValue == "" {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(formValue)

		case reflect.Int:
			value, err := strconv.Atoi(formValue)
			if err != nil {
				return err
			}
			fieldValue.SetInt(int64(value))

		case reflect.Bool:
			value, err := strconv.ParseBool(formValue)
			if err != nil {
				return err
			}
			fieldValue.SetBool(value)

		case reflect.Slice:
			err := setSlice(fieldValue, formValue)
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("not supported type: %v", fieldValue.Kind())
		}
	}
	return nil
}

func findMatchKey(field reflect.StructField) string {
	//first json tag
	jsonKey := field.Tag.Get("json")
	if jsonKey != "" {
		return jsonKey
	}

	//second field name
	//convert to lower-case letter
	return strings.ToLower(field.Name)
}
