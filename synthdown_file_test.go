package synthdown

import (
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
)

func TestDoubledUpJacksError_Error(t *testing.T) {
	err := DoubledUpJacksError{
		errs: []doubledUpJackError{
			{
				module: "square wave",
				jack:   "control",
				positions: []lexer.Position{
					{
						Filename: "synthdown_file_test.go",
						Offset:   0,
						Line:     1,
						Column:   1,
					},
					{
						Filename: "synthdown_file_test.go",
						Offset:   0,
						Line:     5,
						Column:   6,
					},
				},
			},
		},
	}

	expect := `1 error(s):
Jack "control" on Module "square wave" has been patched 2 times: synthdown_file_test.go:1:1, synthdown_file_test.go:5:6`
	received := err.Error()

	if expect != received {
		t.Errorf("expected\n%s\nreceived\n%s\n", expect, received)
	}
}
