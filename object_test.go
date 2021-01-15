package jsontest

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_ParseMap(t *testing.T) {
	j := map[string]interface{}{
		"value": map[string]interface{}{
			"othervalue": "foo",
		},
	}

	obj := ParseMap(j)

	objType := reflect.TypeOf(obj)
	name := objType.Name()

	if name != "Object" {
		t.Errorf("failed to cast map as jsontest.Object, casted to %s", name)
	}
}

func Test_Parse(t *testing.T) {
	m := map[string]interface{}{
		"value": map[string]interface{}{
			"str":    "foo",
			"bool":   true,
			"arr":    []int{1, 2, 3},
			"number": 1,
		},
	}

	j, err := json.Marshal(m)
	if err != nil {
		t.Error("failed to marshal test map into raw JSON")
		return
	}

	obj, err := Parse(j)
	if err != nil {
		t.Errorf("failed to parse valid json payload: %s", err.Error())
		return
	}

	name := reflect.TypeOf(obj).Name()
	if name != "Object" {
		t.Error("failed to cast JSON map to jsontest.Object")
	}
}
