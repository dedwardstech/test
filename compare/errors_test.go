package compare

import (
	"errors"
	"testing"
)

func TestErrors(tt *testing.T) {
	type testArgs struct {
		exp, actual error
	}

	tests := []struct {
		name     string
		args     testArgs
		expected error
	}{
		{
			name: "two nil errors are equal",
			args: testArgs{
				exp:    nil,
				actual: nil,
			},
			expected: nil,
		},
		{
			name: "two of the same errors are equal",
			args: testArgs{
				exp:    errors.New("my cool error"),
				actual: errors.New("my cool error"),
			},
			expected: nil,
		},
		{
			name: "fails when I expect an error but didnt get one",
			args: testArgs{
				exp:    errors.New("really bad error"),
				actual: nil,
			},
			expected: errors.New("wanted err really bad error, but got none"),
		},
		{
			name: "fails when I didnt expect an error but I get one",
			args: testArgs{
				exp:    nil,
				actual: errors.New("oops"),
			},
			expected: errors.New("got unexpected error: oops"),
		},
		{
			name: "fails when I get an error I wasnt expecting",
			args: testArgs{
				exp:    errors.New("bad"),
				actual: errors.New("oh no"),
			},
			expected: errors.New("wanted err bad; got oh no"),
		},
	}

	for _, tc := range tests {
		tt.Run(tc.name, func(t *testing.T) {
			err := Errors(tc.args.exp, tc.args.actual)

			// yes this is a re-implementation of the function
			if tc.expected == nil && err != nil {
				t.Errorf("got unexpected error: %s", err)
				return
			}

			if tc.expected != nil && err == nil {
				t.Errorf("wanted err %v, but got none", tc.expected)
				return
			}

			if tc.expected != nil && (tc.expected.Error() != err.Error()) {
				t.Errorf("wanted err: %v; but got: %v", tc.expected, err)
				return
			}
		})
	}
}
