package jsontest

import (
	"encoding/json"
	"errors"
)

var (
	// ErrPathIndexFailed indicates the path given points to a value that isn't a map
	ErrPathIndexFailed = errors.New("cannot index into non-map type")

	// ErrPropertyDoesNotExist indicates that the path given does not exist in the JSON
	ErrPropertyDoesNotExist = errors.New("json path does not exist")
)

// Object represents a JSON object with some methods
// attached to make it easy to ask questions/make queries
// about the object
type Object map[string]interface{}

// ParseMap casts a map as a jsontest.Object type
func ParseMap(m map[string]interface{}) Object {
	var obj Object = m

	return obj
}

// Parse unmarshals a JSON payload and casts it to an Object
func Parse(payload []byte) (Object, error) {
	var m map[string]interface{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		return nil, err
	}

	return ParseMap(m), nil
}
