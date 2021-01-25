package compare

import "testing"

func Test_Slice(tt *testing.T) {
    type testArgs struct {
        a, b []interface{}
    }

    tests := []struct{
        name string
        args testArgs
        expected bool
    }{
        {
            name: "compares simple slices",
            args: testArgs{
                a: []interface{}{"a", "b", "c" },
                b: []interface{}{"a", "b", "c" },
            },
            expected: true,
        },
        {
            name: "compares slices with the same items in different indexes",
            args: testArgs{
                a: []interface{}{1, 2, 3},
                b: []interface{}{3, 2, 1},
            },
            expected: true,
        },
        {
            name: "compares slices with duplicate items",
            args: testArgs{
                a: []interface{}{1, 2, 2, 3, 3, 3},
                b: []interface{}{3, 3, 3, 2, 2, 1},
            },
            expected: true,
        },
        {
            name: "slices with different lengths are never equal",
            args: testArgs{
                a: []interface{}{1, 2, 3},
                b: []interface{}{1, 2, 3, 4},
            },
            expected: false,
        },
        {
            name: "slices with different elements are not equal",
            args: testArgs{
                a: []interface{}{1, 2, 3},
                b: []interface{}{1, 4, 3},
            },
            expected: false,
        },
        {
            name: "slices with different amounts of the same item are not equal",
            args: testArgs{
                a: []interface{}{1, 2, 2, 3, 3, 3},
                b: []interface{}{1, 2, 2, 3, 3, 3, 3},
            },
            expected: false,
        },
    }

    for _, tc := range tests {
        tt.Run(tc.name, func(t *testing.T) {
            eq := Slice(tc.args.a, tc.args.b)
            if eq != tc.expected {
                if tc.expected {
                    t.Errorf("expected %v and %v to be equal", tc.args.a, tc.args.b)
                } else {
                    t.Errorf("expected %v and %v not to be equal", tc.args.a, tc.args.b)
                }
            }
        })
    }
}

func TestSliceStrict(tt *testing.T) {
    type testArgs struct {
        a, b []interface{}
    }

    tests := []struct{
        name string
        args testArgs
        expected bool
    }{
        {
            name: "slices with the same elements in the same indexes are strictly equal",
            args: testArgs{
                a: []interface{}{1, 2, 3},
                b: []interface{}{1, 2, 3},
            },
            expected: true,
        },
        {
            name: "slices of unequal length are not strictly equal",
            args: testArgs{
                a: []interface{}{1},
                b: []interface{}{1, 2, 3},
            },
            expected: false,
        },
        {
            name: "slices with the same elements in different index are not strictly equal",
            args: testArgs{
                a: []interface{}{1, 2, 3},
                b: []interface{}{3, 2, 1},
            },
            expected: false,
        },
        {
            name: "slices with different elements are not strictly equal",
            args: testArgs{
                a: []interface{}{1, 2, 3},
                b: []interface{}{1, 2, 4},
            },
        },
    }

    for _, tc := range tests {
        tt.Run(tc.name, func(t *testing.T) {
            res := SliceStrict(tc.args.a, tc.args.b)
            if tc.expected != res {
                if tc.expected {
                    t.Errorf("expected %v and %v to be equal", tc.args.a, tc.args.b)
                } else {
                    t.Errorf("expected %v and %v not to be equal", tc.args.a, tc.args.b)
                }
            }
        })
    }
}