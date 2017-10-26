package httpjson

import (
	"net/http"
	"github.com/pkg/errors"
	"reflect"
	"fmt"
	"strconv"
	"strings"
)

var (
	nilError = errors.New("argument nil error")
)

var (
	sliceDelim string = " "
)

func Unmarshal(r *http.Request, v interface{}) error{
	if r == nil || v == nil {
		return nilError
	}

	valSet := reflect.ValueOf(v)

	for i := 0; i < valSet.Type().Elem().NumField(); i++ {
		key := valSet.Elem().Type().Field(i).Tag.Get("json")
		if len(key) == 0 {
			return errors.New(fmt.Sprintf("%s Field has no json tags", valSet.Type().Field(i).Name))
		}
		switch valSet.Elem().Field(i).Kind() {
		case reflect.String:
			valSet.Elem().Field(i).SetString(r.FormValue(key))
			break;

		case reflect.Int:
			value, err := strconv.Atoi(r.FormValue(key))
			if err != nil {
				return err
			}
			valSet.Elem().Field(i).SetInt(int64(value))
			break;

		case reflect.Bool:
			value, err := strconv.ParseBool(r.FormValue(key))
			if err != nil {
				return err
			}
			valSet.Elem().Field(i).SetBool(value)
			break;

		case reflect.Slice:
			value := strings.Fields(r.FormValue(key))
			valSet.Elem().Field(i).Set(reflect.ValueOf(value))
			break;

		default:
			return errors.New(fmt.Sprintln("Not Supported Type ", valSet.Elem().Field(i).Kind()))
			break;
		}
	}
	return nil
}

