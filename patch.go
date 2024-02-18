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

// LastPatchHasOutputError is an error signifying that the
// patch starts with an output, which is a logical absurdity; if it
// has an output from _somewhere_ then it cannot be the start of a patch
// unless it plugs into its self, or is a snippet.
//
// Neither of which are supported, and wont be until we suddenly realise
// it should be
type LastPatchHasOutputError struct {
	m Module
}

// Error satisifies the `error` interface, returning a message explaining
// where the errant initial output lives
func (e LastPatchHasOutputError) Error() string {
	return fmt.Sprintf("the last module of the patch described at %s has an output", e.m.Pos)
}

// MissingInputError is raised where a module should have an input, but
// doesn't
type MissingInputError struct {
	m Module
}

// Error returns an explanation of which module is missing an input jack
func (e MissingInputError) Error() string {
	return fmt.Sprintf("missing input jack at %s", e.m.Pos)
}

// MissingOutputError is raised where a module should have an output, but
// doesn't
type MissingOutputError struct {
	m Module
}

// Error returns an explanation of which module is missing an output jack
func (e MissingOutputError) Error() string {
	return fmt.Sprintf("missing output jack at %s", e.m.Pos)
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

	// Ensure the last module _doesn't_ have an output
	last := p.Modules[modulesLen-1]
	if last.Output != nil {
		return LastPatchHasOutputError{last}
	}

	if modulesLen <= 2 {
		return
	}

	// Ensure every other module has an input and an output
	for i := 1; i < modulesLen-1; i++ {
		mod := p.Modules[i]

		if mod.Input == nil {
			return MissingInputError{mod}
		}

		if mod.Output == nil {
			return MissingOutputError{mod}
		}
	}

	return
}
