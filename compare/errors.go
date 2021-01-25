package compare

import (
	"fmt"

	"errors"
)

// Errors throws an error if the error you expected is different
// from the error you received
//
// There are 5 cases that are considered by this function
//  1. I didn't expect an error and I didn't get one
//  2. I got the error I expected
//  3. I expected an error but didn't get one
//  4. I didn't expect an error but I got one
//  5. I got an error that is different from what I expected
//
//   Errors(nil, nil) => nil
//   Errors("err", "err") => nil
//   Errors("expect an error", nil) => wanted err expected an error, but got none
//   Errors(nil, "shouldn't happen") => got unexpected error: shouldn't happen
//   Errors("err a", "err b") => wanted err err a; got err b
func Errors(expected, actual error) error {

	if expected == nil && actual != nil {
		return errors.New(fmt.Sprintf("got unexpected error: %s", actual.Error()))
	}

	if expected != nil && actual == nil {
		return errors.New(fmt.Sprintf("wanted err %s, but got none", expected.Error()))
	}

	if expected != nil && (expected.Error() != actual.Error()) {
		return errors.New(fmt.Sprintf("wanted err %s; got %s", expected.Error(), actual.Error()))
	}

	return nil
}
