package jsont

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

var (
	// ErrPathIndexFailed indicates the path given points to a value that isn't a map.
	//
	// Given a path "my.cool.path", where "my.cool" is filled by a primitive value.
	// The path exists, but you can't go any further because my.cool isn't a map.
	// Where as ErrPropertyDoesNotExist indicates that the path simply doesn't
	// exist in the payload.
	ErrPathIndexFailed = errors.New("cannot index into non-map type")

	// ErrPropertyDoesNotExist indicates that the path given does not exist in the JSON
	ErrPropertyDoesNotExist = errors.New("json path does not exist")

	// ErrFailedTypeCast indicates you asked for a string, int, array, etc.. but the
	// value in the path was a different type.
	ErrFailedTypeCast = errors.New("")
)

// TypeCastError indicates the value at a property path was not the same
// type of value you wanted.
type TypeCastError struct {
	wantedType, gotType string
}

func (e TypeCastError) Error() string {
	return fmt.Sprintf("attempted to type %s as %s", e.gotType, e.wantedType)
}

// NewTypeCastError creates a TypeCastError given two different types
func NewTypeCastError(wanted, got reflect.Type) TypeCastError {
	return TypeCastError{
		wantedType: wanted.Name(),
		gotType:    got.Name(),
	}
}
