package compare

import (
	"reflect"

	"github.com/dedwardstech/test/internal/count"
)

// Slice compares the two slices given and checks that they have the
// same values in them. It does not care if those values are in the same
// index.
//
//   These two arrays have the same values in different indexes and
//   would not return an error in the Compare function
//   {1, 2, 3, 4}
//   {4, 3, 2, 1}
func Slice(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	cmap := count.SliceItems(a)

	for _, item := range b {
		c, ok := cmap[item]
		if !ok || c == 0 {
			return false
		} else {
			cmap[item]--
		}
	}

	return true
}

// SliceStrict compares two slices by checking that every element in A is the same element in the same
// index
func SliceStrict(a, b []interface{}) bool {
	if len(a) < len(b) {
		return false
	}

	l := len(a)

	for i := 0; i < l; i++ {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}
