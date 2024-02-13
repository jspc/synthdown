package main

import (
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
)

func TestPatch_Errors(t *testing.T) {
	for _, test := range []struct {
		name              string
		p                 Patch
		expectErrorString string
	}{
		{"no modules returns nothing", Patch{}, ""},
		{"first module has an input", Patch{
			Modules: []Module{
				{
					Pos: lexer.Position{
						Filename: "test.sdown",
						Offset:   10,
						Line:     1,
						Column:   1,
					},
					Input: &InputOutput{
						Value: "trigger",
					},
					Name: "flooper",
					Args: []Arg{
						{
							Key:   "F",
							Value: &Value{},
						},
					},
				},
			},
		}, "patch described at test.sdown:1:1 has an input"},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := test.p.Validate()
			if err == nil && test.expectErrorString != "" {
				t.Errorf("expected error, received none")
			} else if err != nil && test.expectErrorString == "" {
				t.Errorf("unexpected error %q", err)
			} else if err != nil && test.expectErrorString != "" {
				rcvd := err.Error()
				if test.expectErrorString != rcvd {
					t.Errorf("expected\n%s\nreceived\n%s", test.expectErrorString, rcvd)
				}
			}
		})
	}
}
