package diff

import (
    "testing"

    "github.com/dedwardstech/test/compare"
)

func Test_SliceDiff_Error(tt *testing.T) {
    tests := []struct {
        name     string
        diff     SliceDiff
        expected string
    }{
        {
            name: "separates diff values with a semi-colon",
            diff: SliceDiff{
                AOnly: []interface{}{"a", "b"},
                BOnly: []interface{}{"c", "d"},
            },
            expected: "values only in A: a, b; values only in B: c, d",
        },
        {
            name: "formats diff for A only correctly",
            diff: SliceDiff{
                AOnly: []interface{}{"a", "b"},
            },
            expected: "values only in A: a, b",
        },
        {
            name: "formats diff for B only correctly",
            diff: SliceDiff{
                BOnly: []interface{}{"c", "d"},
            },
            expected: "values only in B: c, d",
        },
    }

    for _, tc := range tests {
        tt.Run(tc.name, func(t *testing.T) {
            if tc.diff.Error() != tc.expected {
                t.Errorf("expected: %s\ngot: %s", tc.expected, tc.diff.Error())
            }
        })
    }
}

func Test_Slice(tt *testing.T) {
    type testArgs struct {
        a, b []interface{}
    }
    tests := []struct {
        name     string
        args     testArgs
        expected error
    }{
        {
            name: "returns no error if there is no difference between the given slices",
            args: testArgs{
                a: []interface{}{"a", "b"},
                b: []interface{}{"a", "b"},
            },
            expected: nil,
        },
        {
            name: "does not respect element indexes",
            args: testArgs{
                a: []interface{}{"a", "b"},
                b: []interface{}{"b", "a"},
            },
            expected: nil,
        },
    }

    for _, tc := range tests {
        tt.Run(tc.name, func(t *testing.T) {
            err := Slice(tc.args.a, tc.args.b)

            e := compare.Errors(tc.expected, err)
            if e != nil {
                t.Error(e)
                return
            }
        })
    }
}
