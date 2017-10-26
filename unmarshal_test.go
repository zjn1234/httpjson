package httpjson

import (
	"net/http"
	"net/url"
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"

)

type testStruct struct {
	Sstring string 	`json:"sstring"`
	Iint int 		`json:"iint"`
	Bbool bool 		`json:"bbool"`
}

var (
	svalue string = "abc"
	ivalue int = 123
	bvalue bool = true
)

func testInitRequestData() url.Values {
	//make up request
	value := make(url.Values, 10)

	value["sstring"] = []string{svalue}
	value["iint"] = []string{strconv.Itoa(ivalue)}
	value["bbool"] = []string{strconv.FormatBool(bvalue)}

	return value
}

func TestUnmarshal(t *testing.T) {
	m := testInitRequestData()
	res := &http.Request{
		Form:m,
	}

	s := &testStruct{}
	err := Unmarshal(res, s)
	assert.Nil(t, err)
	assert.Equal(t, svalue, s.Sstring)
	assert.Equal(t, ivalue, s.Iint)
	assert.Equal(t, bvalue, s.Bbool)
}

func TestNil(t *testing.T) {
	m := testInitRequestData()
	res := &http.Request{
		Form:m,
	}

	err := Unmarshal(res, nil)
	assert.NotNil(t, err)
}
