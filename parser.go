package synthdown

import (
	"io"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// Module represents a modular synthesiser 'module', such as
// a Square Wave generator, or an Envelope Filter.
//
// Modules contain:
//  1. An input- an audio or control voltage jack (except the first module in a patch which mustn't contain an input)
//  2. A name for the module (so 'Sine', or 'VCA', or whatever)
//  3. Arguments, such as A/D/S/R for an Envelope
//  4. An output jack (except the last module which must omit the final output)
//
// This, effectively, allows us to represent the pertinent information to draw a patch diagram
type Module struct {
	Pos lexer.Position

	Input  *Jack  `parser:"@@?"`
	Name   string `parser:"@Ident"`
	Args   []Arg  `parser:"'[' @@* (',' @@)* ']'"`
	Output *Jack  `parser:"@@?"`
}

// Jack represents a named jack on a synth module, such as the FM jack
// on a wave generator
type Jack struct {
	Pos lexer.Position

	Name string `parser:"'(' @Ident ')'"`
}

// Arg represents the various knobs and dials and twiddles on a synth module, and
// the value it should be set to, such as `Volume: 11` on a mixer module
type Arg struct {
	Pos lexer.Position

	Key   string `parser:"@Ident ':'"`
	Value *Value `parser:"@@"`
}

// New takes a filename containing, hopefully, valid synthdown notation. It returns
// a SynthdownFile type when such a file is found, or one of a series of errors,
// including:
//   - `fs.ErrNotExist` - the specified file does not exist
//   - `*participle.UnexpectedTokenError` - the synthdown notation is invalid
//   - `FirstPatchHasInputError` - the first module in a patch has an input, which is an absurdity
func New(fn string) (p *SynthdownFile, err error) {
	// #nosec: G304
	f, err := os.Open(fn)
	if err != nil {
		return
	}

	defer f.Close()

	return readPatches(fn, f)
}

func readPatches(fn string, input io.Reader) (p *SynthdownFile, err error) {
	parser := participle.MustBuild[SynthdownFile](
		participle.UseLookahead(2),
		participle.Elide("Comment"),
		participle.Unquote(),
	)

	p, err = parser.Parse(fn, input)
	if err != nil {
		return
	}

	return p, p.Validate(nil)
}
