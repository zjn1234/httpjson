package httpjson

import (
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"strings"
)

type testStringStruct struct {
	Sstring string   `json:"sstring"`
	Iint    int      `json:"iint"`
	Bbool   bool     `json:"bbool"`
	Slice   []string `json:"slice"`
}

type testIntStruct struct {
	Sstring string `json:"sstring"`
	Iint    int    `json:"iint"`
	Bbool   bool   `json:"bbool"`
	Slice   []int  `json:"slice"`
}

type testBoolStruct struct {
	Sstring string `json:"sstring"`
	Iint    int    `json:"iint"`
	Bbool   bool   `json:"bbool"`
	Slice   []bool `json:"slice"`
}

type testFloat32Struct struct {
	Sstring string    `json:"sstring"`
	Iint    int       `json:"iint"`
	Bbool   bool      `json:"bbool"`
	Slice   []float32 `json:"slice"`
}

type testFloat64Struct struct {
	Sstring string    `json:"sstring"`
	Iint    int       `json:"iint"`
	Bbool   bool      `json:"bbool"`
	Slice   []float64 `json:"slice"`
}

var (
	svalue             = "abc"
	ivalue             = 123
	bvalue             = true
	Delims             = []rune{' ', ','}
	intSlice           = []int{1, 2, 3}
	stringSlice        = []string{"1", "2", "3"}
	float32StringSlice = []string{"2.5", "3.6", "10.8"}
	float32Slice       = []float32{2.5, 3.6, 10.8}
	boolStringSlice    = []string{"false", "true", "false"}
	boolSlice          = []bool{false, true, false}
	float64StringSlice = []string{"2.5555555", "3.66666666", "10.888888888"}
	float64Slice       = []float64{2.5555555, 3.66666666, 10.888888888}
)

func testStringSlice(t *testing.T, delim rune, r *http.Request) {
	s := &testStringStruct{}

	SetDelim(delim)
	r.Form["slice"] = []string{strings.Join(stringSlice, string(delim))}

	//unmarshal
	err := Unmarshal(r, s)
	assert.Nil(t, err)
	assert.Equal(t, svalue, s.Sstring)
	assert.Equal(t, ivalue, s.Iint)
	assert.Equal(t, bvalue, s.Bbool)
	assert.Equal(t, stringSlice, s.Slice)
}

func testFloat32Slice(t *testing.T, delim rune, r *http.Request) {
	s := &testFloat32Struct{}

	SetDelim(delim)
	r.Form["slice"] = []string{strings.Join(float32StringSlice, string(delim))}

	//unmarshal
	err := Unmarshal(r, s)
	assert.Nil(t, err)
	assert.Equal(t, svalue, s.Sstring)
	assert.Equal(t, ivalue, s.Iint)
	assert.Equal(t, bvalue, s.Bbool)
	assert.Equal(t, float32Slice, s.Slice)

}

func testFloat64Slice(t *testing.T, delim rune, r *http.Request) {
	s := &testFloat64Struct{}

	SetDelim(delim)
	r.Form["slice"] = []string{strings.Join(float64StringSlice, string(delim))}

	//unmarshal
	err := Unmarshal(r, s)
	assert.Nil(t, err)
	assert.Equal(t, svalue, s.Sstring)
	assert.Equal(t, ivalue, s.Iint)
	assert.Equal(t, bvalue, s.Bbool)
	assert.Equal(t, float64Slice, s.Slice)

}

func testIntSlice(t *testing.T, delim rune, r *http.Request) {
	s := &testIntStruct{}

	SetDelim(delim)
	r.Form["slice"] = []string{strings.Join(stringSlice, string(delim))}

	//unmarshal
	err := Unmarshal(r, s)
	assert.Nil(t, err)
	assert.Equal(t, svalue, s.Sstring)
	assert.Equal(t, ivalue, s.Iint)
	assert.Equal(t, bvalue, s.Bbool)
	assert.Equal(t, intSlice, s.Slice)

}

func testBoolSlice(t *testing.T, delim rune, r *http.Request) {
	s := &testBoolStruct{}
	SetDelim(delim)
	r.Form["slice"] = []string{strings.Join(boolStringSlice, string(delim))}

	//unmarshal
	err := Unmarshal(r, s)
	assert.Nil(t, err)
	assert.Equal(t, svalue, s.Sstring)
	assert.Equal(t, ivalue, s.Iint)
	assert.Equal(t, bvalue, s.Bbool)
	assert.Equal(t, boolSlice, s.Slice)

}

func initData() *http.Request {
	value := make(url.Values, 10)
	value["sstring"] = []string{svalue}
	value["iint"] = []string{strconv.Itoa(ivalue)}
	value["bbool"] = []string{strconv.FormatBool(bvalue)}

	return &http.Request{
		Form: value,
	}
}

func TestAllData(t *testing.T) {
	request := initData()
	for _, delim := range Delims {
		testBoolSlice(t, delim, request)
		testStringSlice(t, delim, request)
		testIntSlice(t, delim, request)

		testFloat32Slice(t, delim, request)
		testFloat64Slice(t, delim, request)
	}
}
