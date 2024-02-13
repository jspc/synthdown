package synthdown

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

// FirstPatchHasInputError is an error signifying that the
// patch starts with an input, which is a logical absurdity; if it
// has an input from _somewhere_ then it cannot be the start of a patch
// unless it plugs into its self, or is a snippet.
//
// Neither of which are supported, and wont be until we suddenly realise
// it should be
type FirstPatchHasInputError struct {
	m Module
}

// Error satisifies the `error` interface, returning a message explaining
// where the errant initial input lives
func (e FirstPatchHasInputError) Error() string {
	return fmt.Sprintf("the first module of the patch described at %s has an input", e.m.Pos)
}

// Patch represents a set of modules cabled together from jack to jack
// which hopefully makes an interesting noise
type Patch struct {
	Pos lexer.Position

	Modules []Module `parser:"@@ ('-' '>' @@)*"`
}

// Validate runs through a patch and checks it against a set of rules
func (p Patch) Validate() (err error) {
	modulesLen := len(p.Modules)

	if modulesLen == 0 {
		return
	}

	// Ensure first module _doesn't_ have an input
	first := p.Modules[0]
	if first.Input != nil {
		return FirstPatchHasInputError{first}
	}

	return
}
