package synthdown

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
					Input: &Jack{
						Name: "trigger",
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
		}, "the first module of the patch described at test.sdown:1:1 has an input"},
		{"last module has an input", Patch{
			Modules: []Module{
				{
					Pos: lexer.Position{
						Filename: "test.sdown",
						Offset:   10,
						Line:     1,
						Column:   1,
					},
					Output: &Jack{
						Name: "trigger",
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
		}, "the last module of the patch described at test.sdown:1:1 has an output"},
		{"midde module is missing an input", Patch{
			Modules: []Module{
				{
					Output: &Jack{
						Name: "trigger",
					},
					Name: "flooper",
					Args: []Arg{
						{
							Key:   "F",
							Value: &Value{},
						},
					},
				},
				{
					Pos: lexer.Position{
						Filename: "test.sdown",
						Offset:   10,
						Line:     1,
						Column:   1,
					},
					Output: &Jack{
						Name: "cv1",
					},
					Name: "fridgefreezer",
					Args: []Arg{
						{
							Key:   "Ytho",
							Value: &Value{},
						},
					},
				},
				{
					Input: &Jack{
						Name: "in",
					},
					Name: "flimflammer",
					Args: []Arg{
						{
							Key:   "X",
							Value: &Value{},
						},
					},
				},
			},
		}, "missing input jack at test.sdown:1:1"},
		{"midde module is missing an output", Patch{
			Modules: []Module{
				{
					Output: &Jack{
						Name: "trigger",
					},
					Name: "flooper",
					Args: []Arg{
						{
							Key:   "F",
							Value: &Value{},
						},
					},
				},
				{
					Pos: lexer.Position{
						Filename: "test.sdown",
						Offset:   10,
						Line:     1,
						Column:   1,
					},
					Input: &Jack{
						Name: "cv1",
					},
					Name: "fridgefreezer",
					Args: []Arg{
						{
							Key:   "Ytho",
							Value: &Value{},
						},
					},
				},
				{
					Input: &Jack{
						Name: "in",
					},
					Name: "flimflammer",
					Args: []Arg{
						{
							Key:   "X",
							Value: &Value{},
						},
					},
				},
			},
		}, "missing output jack at test.sdown:1:1"},
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
