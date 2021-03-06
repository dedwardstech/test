// Package jsont provides types and functions to make testing
// JSON payloads easier
package jsont

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	strType     = reflect.TypeOf("")
	boolType    = reflect.TypeOf(true)
	intType     = reflect.TypeOf(1)
	int64Type   = reflect.TypeOf(int64(1))
	float64Type = reflect.TypeOf(1.0)
	objType     = reflect.TypeOf(Object{})
	sliceType   = reflect.TypeOf([]interface{}{})
)

// Object represents a JSON object with some methods
// attached to make it easy to ask questions/make queries
// about the object.
type Object map[string]interface{}

// Set casts a map as a Object type
func Set(m map[string]interface{}) Object {
	var obj Object = m

	return obj
}

// Unmarshal unmarshals a JSON payload and casts it to an Object
func Unmarshal(payload []byte) (Object, error) {
	var m map[string]interface{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		return nil, err
	}

	return Set(m), nil
}

// Has is used when you ask the question, "Does this property path exist?"
func (o *Object) Has(propertyPath string) (bool, error) {
	_, err := parsePathValue(*o, propertyPath)
	if err != nil {
		return false, err
	}

	return true, nil
}

// HasKeys test whether the object has the given key paths available
func (o Object) HasKeys(propertyPaths []string) (bool, error) {
	errs := make([]string, 0)

	for _, path := range propertyPaths {
		present, err := o.Has(path)
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", path, err))
			continue
		}

		if !present {
			errs = append(errs, fmt.Sprintf("%s: json path does not exist", path))
			continue
		}
	}

	if len(errs) > 0 {
		errMsg := strings.Join(errs, "\n")
		return false, errors.New(errMsg)
	}

	return true, nil
}

// Get is used to when you ask the question, "What value is associated with this property path?"
func (o Object) Get(propertyPath string) (interface{}, error) {
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// GetStr is used to extract a string value from an object
func (o Object) GetStr(propertyPath string) (string, error) {
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return "", err
	}

	str, ok := val.(string)
	if !ok {
		return "", NewTypeCastError(strType, reflect.TypeOf(val))
	}

	return str, nil
}

// GetNumber is used to extract a number value, as a float64, from an object
func (o Object) GetNumber(propertyPath string) (float64, error) {
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return float64(-1), err
	}

	num, ok := val.(float64)
	if !ok {
		return float64(-1), NewTypeCastError(float64Type, reflect.TypeOf(val))
	}

	return num, nil
}

// GetInt64 is used to extract a number value, as a int64, from an object
func (o Object) GetInt64(propertyPath string) (int64, error) {
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return int64(-1), err
	}

	num, ok := val.(int64)
	if !ok {
		return int64(-1), NewTypeCastError(int64Type, reflect.TypeOf(val))
	}

	return num, nil
}

// GetBool is used to extract a key whose value is an bool
func (o Object) GetBool(propertyPath string) (bool, error) {
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return false, err
	}

	b, ok := val.(bool)
	if !ok {
		return false, NewTypeCastError(boolType, reflect.TypeOf(val))
	}

	return b, nil
}

// GetSlice extracts a slice from an object
func (o Object) GetSlice(propertyPath string) ([]interface{}, error) {
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return nil, err
	}

	sl, ok := val.([]interface{})
	if !ok {
		return nil, NewTypeCastError(sliceType, reflect.TypeOf(val))
	}

	return sl, nil
}

// GetObj is used to extract a key whose value is an object
func (o Object) GetObj(propertyPath string) (Object, error) {
	fmt.Println("getting value", propertyPath)
	val, err := parsePathValue(o, propertyPath)
	if err != nil {
		return nil, err
	}

	fmt.Printf("got %v\n", val)
	m, ok := val.(map[string]interface{})
	if !ok {
		return nil, NewTypeCastError(objType, reflect.TypeOf(val))
	}

	var obj Object = m
	return obj, nil
}
