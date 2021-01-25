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

func TestObject_Get(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   interface{}
		err        error
	}{
		{
			name: "gets an interface value from a simple Object",
			path: "foo",
			obj: Object{
				"foo": "bar",
			},
			expected: "bar",
			err:      nil,
		},
		{
			name: "gets an interface value from a nested Object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": 1.0,
					},
				},
			},
			expected: 1.0,
			err:      nil,
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: nil,
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": 1.0,
					},
				},
			},
			expected: nil,
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {
			val, err := tc.obj.Get(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if val != tc.expected {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestObject_GetStr(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   string
		err        error
	}{
		{
			name: "gets a string value from a simple object",
			path: "foo",
			obj: Object{
				"foo": "str",
			},
			expected: "str",
			err:      nil,
		},
		{
			name: "gets a string value from a nested object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": "str",
					},
				},
			},
			expected: "str",
			err:      nil,
		},
		{
			name: "throws an error if value isnt a string",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": 1,
					},
				},
			},
			expected: "",
			err:      NewTypeCastError(strType, intType),
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: "",
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": "str",
					},
				},
			},
			expected: "",
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {
			val, err := tc.obj.GetStr(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if val != tc.expected {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestObject_GetNumber(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   float64
		err        error
	}{
		{
			name: "gets a float64 value from a simple object",
			path: "foo",
			obj: Object{
				"foo": float64(1),
			},
			expected: float64(1),
			err:      nil,
		},
		{
			name: "gets a float64 value from a nested object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": float64(2),
					},
				},
			},
			expected: float64(2),
			err:      nil,
		},
		{
			name: "throws an error if value isnt a number",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": "str",
					},
				},
			},
			expected: float64(-1),
			err:      NewTypeCastError(float64Type, strType),
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: float64(-1),
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": "str",
					},
				},
			},
			expected: float64(-1),
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {
			val, err := tc.obj.GetNumber(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if val != tc.expected {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestObject_GetInt64(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   int64
		err        error
	}{
		{
			name: "gets a string value from a simple object",
			path: "foo",
			obj: Object{
				"foo": int64(1000),
			},
			expected: int64(1000),
			err:      nil,
		},
		{
			name: "gets a string value from a nested object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": int64(1000),
					},
				},
			},
			expected: int64(1000),
			err:      nil,
		},
		{
			name: "throws an error if value isnt an int64",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": float64(1000),
					},
				},
			},
			expected: int64(-1),
			err:      NewTypeCastError(int64Type, float64Type),
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: int64(-1),
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": int64(1000),
					},
				},
			},
			expected: int64(-1),
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {
			val, err := tc.obj.GetInt64(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if val != tc.expected {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestObject_GetBool(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   bool
		err        error
	}{
		{
			name: "gets a string value from a simple object",
			path: "foo",
			obj: Object{
				"foo": false,
			},
			expected: false,
			err:      nil,
		},
		{
			name: "gets a string value from a nested object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": false,
					},
				},
			},
			expected: false,
			err:      nil,
		},
		{
			name: "throws an error if value isnt an int64",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": 1,
					},
				},
			},
			expected: false,
			err:      NewTypeCastError(boolType, intType),
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: false,
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": false,
					},
				},
			},
			expected: false,
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {
			val, err := tc.obj.GetBool(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if val != tc.expected {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestObject_GetSlice(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   []interface{}
		err        error
	}{
		{
			name: "gets a slice value from a simple object",
			path: "foo",
			obj: Object{
				"foo": []interface{}{1, 2, 3},
			},
			expected: []interface{}{1, 2, 3},
			err:      nil,
		},
		{
			name: "gets a slice value from a nested object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": []interface{}{1, 2, 3},
					},
				},
			},
			expected: []interface{}{1, 2, 3},
			err:      nil,
		},
		{
			name: "throws an error if value isnt a slice",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": "str",
					},
				},
			},
			expected: nil,
			err:      NewTypeCastError(sliceType, strType),
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": []interface{}{1, 2, 3},
				},
			},
			expected: nil,
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": []interface{}{1, 2, 3},
					},
				},
			},
			expected: nil,
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {
			val, err := tc.obj.GetSlice(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if !compare.SliceStrict(val, tc.expected) {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestObject_GetObj(tt *testing.T) {
	testcases := []struct {
		name, path string
		obj        Object
		expected   Object
		err        error
	}{
		{
			name: "gets a string value from a simple object",
			path: "foo",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: Object{"bar": true},
			err:      nil,
		},
		{
			name: "gets a string value from a nested object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": map[string]interface{}{
							"something": true,
						},
					},
				},
			},
			expected: Object{"something": true},
			err:      nil,
		},
		{
			name: "throws an error if value isnt an Object",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": float64(1000),
					},
				},
			},
			expected: nil,
			err:      NewTypeCastError(objType, float64Type),
		},
		{
			name: "throws an error if the path exists but you cant index into it",
			path: "foo.bar.baz",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": true,
				},
			},
			expected: nil,
			err:      ErrPathIndexFailed,
		},
		{
			name: "throws an error if the path does not exist",
			path: "foo.bar.foobar",
			obj: Object{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"baz": map[string]interface{}{
							"something": true,
						},
					},
				},
			},
			expected: nil,
			err:      ErrPropertyDoesNotExist,
		},
	}

	for _, tc := range testcases {
		tt.Run(tc.name, func(t *testing.T) {

			val, err := tc.obj.GetObj(tc.path)
			testErr := compare.Errors(tc.err, err)
			if testErr != nil {
				t.Error(testErr)
				return
			}

			if !reflect.DeepEqual(val, tc.expected) {
				t.Errorf("expected value: %v, got: %v", tc.expected, val)
			}
		})
	}
}
