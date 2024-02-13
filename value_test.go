package synthdown

import (
	"testing"
)

var (
	foo                  = "foo"
	onePointNineNineNine = 1.999
	one                  = 1
	twoPointFive         = "2.5"
)

func TestValue_Float64(t *testing.T) {
	for _, test := range []struct {
		name        string
		v           Value
		expect      float64
		expectPanic bool
	}{
		{"nil value returns 0", Value{}, 0, false},
		{"non-float value as string panics", Value{Str: &foo}, 0, true},
		{"float value returns exactly", Value{Float: &onePointNineNineNine}, 1.999, false},
		{"int value returns as float", Value{Int: &one}, 1.0, false},
		{"valid string value returns as float", Value{Str: &twoPointFive}, 2.5, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err == nil && test.expectPanic {
					t.Errorf("expected error, received none")
				} else if err != nil && !test.expectPanic {
					t.Errorf("unexpected error %#v", err)
				}
			}()

			rcvd := test.v.Float64()
			if test.expect != rcvd {
				t.Errorf("expected %f, received %f", test.expect, rcvd)
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	for _, test := range []struct {
		name   string
		v      Value
		expect string
	}{
		{"nil value returns 0", Value{}, "0"},
		{"float value returns as string", Value{Float: &onePointNineNineNine}, "1.999"},
		{"int value returns as float", Value{Int: &one}, "1"},
		{"valid string value returns exactly", Value{Str: &twoPointFive}, "2.5"},
	} {
		t.Run(test.name, func(t *testing.T) {
			rcvd := test.v.String()
			if test.expect != rcvd {
				t.Errorf("expected %q, received %q", test.expect, rcvd)
			}
		})
	}
}
