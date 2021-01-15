package jsontest

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	// ErrPathIndexFailed indicates the path given points to a value that isn't a map.
	//
	// Given a path "my.cool.path", where "my.cool" is filled by a primitive value.
	// The path exists, but you can't go any further because my.cool isn't a map.
	// Where as ErrPropertyDoesNotExist indicates that the path simply doesn't
	// exist in the payload.
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

// Has is used when you ask the question, "Does this property path exist?"
func (o *Object) Has(propertyPath string) (bool, error) {
	_, err := parsePathValue(*o, propertyPath)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Get is used to when you ask the question, "What value is associated with this property path?"
func (o *Object) Get(propertyPath string) (interface{}, error) {
	val, err := parsePathValue(*o, propertyPath)
	if err != nil {
		return nil, err
	}

	return val, nil
}
