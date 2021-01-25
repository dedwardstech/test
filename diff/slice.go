package diff

import (
    "fmt"
    "strings"

    "github.com/dedwardstech/test/internal/count"
)

// SliceDiff is used to hold the results of diffing two slices.
type SliceDiff struct {
    AOnly []interface{}
    BOnly []interface{}
}

func (d SliceDiff) Error() string {
    aStr := ""
    bStr := ""
    if len(d.AOnly) > 0 {
        diffStr := formatSliceStr(d.AOnly)
        aStr += "values only in A: " + diffStr
    }

    if len(d.BOnly) > 0 {
        diffStr := formatSliceStr(d.BOnly)
        bStr += "values only in B: " + diffStr
    }

    if len(aStr) > 0 && len(bStr) > 0 {
        return aStr + "; " + bStr
    }

    return aStr + bStr
}

func formatSliceStr(s []interface{}) string {
    if len(s) == 0 {
        return ""
    }

    b := make([]string, len(s))

    for i, v := range s {
        b[i] = fmt.Sprintf("%v", v)
    }

    return strings.Join(b, ", ")
}

// Slice calculates the differences between two slices. It finds all of the values that are in A but not in B, and
// it finds all of the values that are in B but not in A.
// The result is implemented as an error because this method is intended to be used in tests, and when testing,
// for the most part, it doesn't matter what the differences are. Only that there are differences and those
// differences can be written to StdOut.
//
// Slices are compared without regard to what indexes the elements are in.
//   a := {1, 2, 3}
//   b := {3, 2, 1}
//
//   diff(a, b) = {}
//
// Duplicate elements are also accounted for.
//   a := {1, 2, 2, 3, 3, 3}
//   b := {1, 2, 3}
//
//   diff(a, b) = {2, 3, 3}
func Slice(a, b []interface{}) error {
    aOnly := make([]interface{}, 0)
    bOnly := make([]interface{}, 0)
    bCount := count.SliceItems(b)

    for _, elem := range a {
        if c, ok := bCount[elem]; !ok || c == 0 {
            aOnly = append(aOnly, elem)

            if c == 0 {
                delete(bCount, elem)
            }
        } else {
            bCount[elem]--
        }
    }

    for key, c := range bCount {
        if c > 0 {
            for i := 0; i <= c; i++ {
                bOnly = append(bOnly, key)
            }
        }
    }

    aOnlyLen, bOnlyLen := len(aOnly), len(bOnly)
    if aOnlyLen > 0 || bOnlyLen > 0 {
        err := SliceDiff{}

        if aOnlyLen > 0 {
            err.AOnly = aOnly
        }

        if bOnlyLen > 0 {
            err.BOnly = bOnly
        }

        return err
    }

    return nil
}
