package jsont

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/dedwardstech/test/compare"
)

func Test_Set(t *testing.T) {
	j := map[string]interface{}{
		"value": map[string]interface{}{
			"othervalue": "foo",
		},
	}

	obj := Set(j)

	objType := reflect.TypeOf(obj)
	name := objType.Name()

	if name != "Object" {
		t.Errorf("failed to cast map as jsont.Object, casted to %s", name)
	}
}

func Test_Unmarshal(t *testing.T) {
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

	obj, err := Unmarshal(j)
	if err != nil {
		t.Errorf("failed to parse valid json payload: %s", err.Error())
		return
	}

	name := reflect.TypeOf(obj).Name()
	if name != "Object" {
		t.Error("failed to cast JSON map to test.Object")
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

	for _, tc := range tests {
		tt.Run(tc.name, func(t *testing.T) {
			res, err := tc.obj.Has(tc.path)
			if e := compare.Errors(tc.err, err); e != nil {
				t.Error(e)
				return
			}

			if res != tc.expected {
				t.Errorf("wanted result %v; got %v", tc.expected, err)
				return
			}
		})
	}
}

//func Test_Object_Get(tt *testing.T) {
//	tests := []struct {
//		name     string
//		obj      Object
//		path     string
//		expected interface{}
//		err      error
//	}{
//		{
//			name: "retrieves value from a simple path",
//			obj: Object{
//				"value": "foo",
//			},
//			path:     "value",
//			expected: "foo",
//			err:      nil,
//		},
//		{
//			name: "retrieves value from a nested path",
//			obj: Object{
//				"value": map[string]interface{}{
//					"othervalue": map[string]interface{}{
//						"end": "foo",
//					},
//				},
//			},
//			path:     "value.othervalue.end",
//			expected: "foo",
//			err:      nil,
//		},
//		{
//			name: "retrieves object values from a nested path",
//			obj: Object{
//				"value": map[string]interface{}{
//					"othervalue": map[string]interface{}{
//						"end": "foo",
//					},
//				},
//			},
//			path: "value.othervalue",
//			expected: map[string]interface{}{
//				"end": "foo",
//			},
//			err: nil,
//		},
//		{
//			name: "throws an error if the property path tries index into a non-map type",
//			obj: Object{
//				"value": map[string]interface{}{
//					"othervalue": "foo",
//				},
//			},
//			path:     "value.othervalue.end",
//			expected: nil,
//			err:      ErrPathIndexFailed,
//		},
//		{
//			name: "throws an error if the property path does not exist",
//			obj: Object{
//				"value": map[string]interface{}{
//					"othervalue": map[string]interface{}{
//						"end": "foo",
//					},
//				},
//			},
//			path:     "value.otherothervalue.end",
//			expected: nil,
//			err:      ErrPropertyDoesNotExist,
//		},
//	}
//
//	for _, testcase := range tests {
//		tt.Run(testcase.name, func(t *testing.T) {
//			var v interface{}
//			err := testcase.obj.Get(testcase.path, &v)
//			if e := test.CompareErrors(testcase.err, err); e != nil {
//				t.Error(e)
//				return
//			}
//
//			if !reflect.DeepEqual(testcase.expected, v) {
//				t.Errorf("wanted result %v; got %v", testcase.expected, v)
//			}
//		})
//	}
//}

//func Test_Object_Keys(tt *testing.T) {
//	tests := []struct {
//		name     string
//		obj      Object
//		expected []string
//	}{
//		{
//			name: "returns keys for a simple object",
//			obj: Object{
//				"key1": true,
//				"key2": "foo",
//				"key3": 1,
//				"key4": nil,
//			},
//			expected: []string{"key1", "key2", "key3", "key4"},
//		},
//	}
//
//	for _, tc := range tests {
//		tt.Run(tc.name, func(t *testing.T) {
//			res := tc.obj.Keys()
//			rSize, eSize := len(res), len(tc.expected)
//			if rSize != eSize {
//				t.Errorf("expected keys to be %v: got %v", tc.expected, res)
//				return
//			}
//
//			i := 0
//			for i < eSize {
//				if res[i] != tc.expected[i] {
//					t.Errorf("expected keys to be %v: got %v", tc.expected, res)
//				}
//				i++
//			}
//		})
//	}
//}
