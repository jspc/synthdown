package synthdown

import (
	"bytes"
	"errors"
	"testing"

	"github.com/alecthomas/participle/v2"
)

var (
	validPatches = []string{
		`square[](out) -> (c1)mixer[c1:10, out:10](out) -> (in)speaker[];`,
		`noise[](white) -> (in)speaker[];`,
		`noise[](white)
-> (in)rand[L:10, R:10](out)
-> (in)vcs[L:10];`,
		`square[](out) -> (c1)mixer[c1:10, out:10](out) -> (in)speaker[];
saw[](out) -> (c2)mixer[c2:10];
sine[](out) -> (c3)mixer[c3:7.5];
`,
		`sequencer[bpm:90](output) -> (trigger)envelope[A:0,D:2,S:0,R:2](out) -> (control)sine[L:10,T:1.5];`, // actual bass drum-alike we use in a simple patch
		``,
		`square[];`,
	}
)

func TestNewPatch_Valid(t *testing.T) {
	for _, p := range validPatches {
		t.Run("", func(t *testing.T) {
			_, err := readPatches("", bytes.NewBufferString(p))
			if err != nil {
				t.Errorf("unexpected error: %#v", err)
			}
		})

	}
}

func TestNew(t *testing.T) {
	for _, test := range []struct {
		fn          string
		expectError bool
	}{
		{"testdata/complex.sdown", false},
		{"testdata/nonsuch", true},
	} {
		t.Run(test.fn, func(t *testing.T) {
			_, err := New(test.fn)
			if err == nil && test.expectError {
				t.Errorf("expected error, received none")
			} else if err != nil && !test.expectError {
				t.Errorf("unexpected error %#v", err)
			}
		})
	}
}

func TestNewPatch_Errors(t *testing.T) {
	var (
		ute  *participle.UnexpectedTokenError
		fphi FirstPatchHasInputError
		lpho LastPatchHasOutputError
		mie  MissingInputError
		moe  MissingOutputError
		duj  DoubledUpJacksError
	)

	for _, test := range []struct {
		name string
		p    string
		errT any
	}{
		{"semicolon only is invalid", ";", &ute},
		{"missing semicolon is invalid", "square[]", &ute},
		{"first module shouldn't have an input", "(in)square[](out);", &fphi},
		{"last module shouldn't have an output", "square[](out);", &lpho},
		{"intermediate modules should have inputs", "square[](out) -> vca[](out) -> (in)mixer[];", &mie},
		{"intermediate modules should have outputs", "square[](out) -> (in)vca[] -> (in)mixer[];", &moe},
		{"jacks mustn't be reused", `square[](out) -> (in)mixer[];
square[](out) -> (left)speaker[];`, &duj},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := readPatches("", bytes.NewBufferString(test.p))
			if err == nil {
				t.Error("expected error, received none")
			}

			if !errors.As(err, test.errT) {
				t.Errorf("expected error of type %T, recveived %T", test.errT, err)

				t.Logf("%#v", err)
			}
		})
	}
}
