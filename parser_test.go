package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/alecthomas/participle/v2"
)

var (
	validPatches = []string{
		`square[](out) -> (c1)mixer[c1:10, out:10](out) -> (in)speaker[](out);`,
		`noise[](white) -> (in)speaker[](out);`,
		`noise[](white)
-> (in)rand[L:10, R:10](out)
-> (in)vcs[L:10](out);`,
		`square[](out) -> (c1)mixer[c1:10, out:10](out) -> (in)speaker[](out);
saw[](out) -> (c2)mixer[c2:10](out);
sine[](out) -> (c3)mixer[c3:7.5](out);
`,
		`sequencer[bpm:90](output) -> (trigger)envelope[A:0,D:2,S:0,R:2](out) -> (control)sine[L:10,T:1.5](out);`, // actual bass drum-alike we use in a simple patch
		``,
		`square[](out);`,
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
		fphi errFirstPatchHasInput
	)

	for _, test := range []struct {
		name string
		p    string
		errT any
	}{
		{"semicolon only is invalid", ";", &ute},
		{"missing semicolon is invalid", "square[](out)", &ute},
		{"first module shouldn't have an input", "(in)square[](out);", &fphi},
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
