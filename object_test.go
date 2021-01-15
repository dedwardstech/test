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

func Test_Object_Has(tt *testing.T) {
	tests := []struct {
		name     string
		obj      Object
		path     string
		expected bool
		err      error
	}{
		{
			name: "handles simple property paths",
			obj: Object{
				"value": map[string]interface{}{
					"othervalue": map[string]interface{}{
						"final": "foo",
					},
				},
			},
			path:     "value",
			expected: true,
			err:      nil,
		},
		{
			name: "handles nested property paths",
			obj: Object{
				"value": map[string]interface{}{
					"othervalue": map[string]interface{}{
						"final": "foo",
					},
				},
			},
			path:     "value.othervalue.final",
			expected: true,
			err:      nil,
		},
		{
			name: "returns an error if the property path does not exist",
			obj: Object{
				"value": map[string]interface{}{
					"othervalue": map[string]interface{}{
						"final": "foo",
					},
				},
			},
			path:     "value.missing.val",
			expected: false,
			err:      ErrPropertyDoesNotExist,
		},
		{
			name: "returns an error if the property path points to a non-map type",
			obj: Object{
				"value": true,
			},
			path:     "value.othervalue",
			expected: false,
			err:      ErrPathIndexFailed,
		},
	}

	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			res, err := test.obj.Has(test.path)
			if test.err == nil && err != nil {
				t.Errorf("got unexpected err: %s", err.Error())
				return
			}

			if test.err != nil && err == nil {
				t.Errorf("wanted err %s, but got none", test.err.Error())
				return
			}

			if test.err != nil && (test.err.Error() != err.Error()) {
				t.Errorf("wanted err %s; got %s", test.err.Error(), err.Error())
				return
			}

			if res != test.expected {
				t.Errorf("wanted result %v; got %v", test.expected, err)
				return
			}
		})
	}
}
