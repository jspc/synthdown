package main

import (
	"io"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Patches struct {
	Pos lexer.Position

	Patches []Patch `parser:"( @@ ';' )*"`
}

type Module struct {
	Pos lexer.Position

	Input  *InputOutput `parser:"@@?"`
	Name   string       `parser:"@Ident"`
	Args   []Arg        `parser:"'[' @@* (',' @@)* ']'"`
	Output InputOutput  `parser:"@@"`
}

type InputOutput struct {
	Pos lexer.Position

	Value string `parser:"'(' @Ident ')'"`
}

type Arg struct {
	Pos lexer.Position

	Key   string `parser:"@Ident ':'"`
	Value *Value `parser:"@@"`
}

func New(fn string) (p *Patches, err error) {
	f, err := os.Open(fn)
	if err != nil {
		return
	}

	defer f.Close()

	return readPatches(fn, f)
}

func readPatches(fn string, input io.Reader) (p *Patches, err error) {
	parser := participle.MustBuild[Patches](
		participle.UseLookahead(2),
		participle.Elide("Comment"),
		participle.Unquote(),
	)

	p, err = parser.Parse(fn, input)
	if err != nil {
		return
	}

	for _, patch := range p.Patches {
		err = patch.Validate()
		if err != nil {
			return
		}
	}

	return
}
