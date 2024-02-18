package synthdown

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

// DoubledUpJacksError returns when one or more jacks exist in many patches
// yet shouldn't.
//
// This error may be incorrect for a number of reasons, including:
//
//  1. Not all jacks on a module have a unique name, such as the sequencer on a POM-400
//  2. You're using a stacking jack, or a splitter, or a multiplier
//
// In these instances the solution is to:
//
//  1. Artificially number each jack, like 'output1', 'output2'
//  2. Set `StackedPatches: true` in the synthdown.ValidationConfiguration struct
//
// passed to (SynthdownFile).Validate()
type DoubledUpJacksError struct {
	errs []doubledUpJackError
}

// Error returns a message indicating the number of errors found, plus the location
// of each individual error
func (e DoubledUpJacksError) Error() string {
	errStrings := make([]string, len(e.errs))
	for i, err := range e.errs {
		errStrings[i] = err.Error()
	}

	return fmt.Sprintf("%d error(s):\n%s", len(errStrings), strings.Join(errStrings, "\n"))
}

// doubledUpJackError returns when a specific jack on a module is
// used more than once in a set of patches, and the ValidationConfiguration
// passed to SynthdownFile.Validation doesn't have StackedPatches set to true.
//
// Default behvaiour is to assume that modules only have one jack by name
type doubledUpJackError struct {
	module, jack string
	positions    []lexer.Position
}

// Error returns a string explaining where a named jack has been used more than once
func (e doubledUpJackError) Error() string {
	positions := make([]string, len(e.positions))
	for i, pos := range e.positions {
		positions[i] = pos.String()
	}

	return fmt.Sprintf("Jack %q on Module %q has been patched %d times: %s",
		e.jack, e.module, len(positions), strings.Join(positions, ", "),
	)
}

// moduleJackMapping tracks which named jacks are in use where, in order to validate
// when jacks are used more than once
type moduleJackMapping map[string]map[string][]lexer.Position

// addMapping takes a module, a jack, and a position in a synthdown file
func (m moduleJackMapping) addMapping(module, jack string, pos lexer.Position) {
	if _, ok := m[module]; !ok {
		m[module] = make(map[string][]lexer.Position)
	}

	if _, ok := m[module][jack]; !ok {
		m[module][jack] = make([]lexer.Position, 0)
	}

	m[module][jack] = append(m[module][jack], pos)
}

// errors combines every doubled-up jack into a single error for convenience
func (m moduleJackMapping) errors() error {
	errs := make([]doubledUpJackError, 0)

	for module, jacks := range m {
		for jack, positions := range jacks {
			if len(positions) > 1 {
				errs = append(errs, doubledUpJackError{module, jack, positions})
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return DoubledUpJacksError{errs}
}

// ValidationConfiguration holds options for enabling/ disabling
// certain validation options.
//
// The default behaviour of any validation is to enable _every_
// validation, which is why this struct is only ever passed as a
// reference.
type ValidationConfiguration struct {
	// StackedPatches implies the use of multipliers/ splitter
	// cables/ stacking plugs which allows us to, effectively,
	// reuse input/output jacks on a module
	StackedPatches bool
}

// SynthdownFile contains a list of Patches, which describe how
// modular synth modules are wired to one another.
type SynthdownFile struct {
	Pos lexer.Position

	Patches []Patch `parser:"( @@ ';' )*"`
}

// Validate runs through a fully read synthdown file in order to:
//
//  1. Run Patch.Validate() on each patch
//  2. Ensure inputs and outputs aren't reused
//
// Note: checking input/ output reuse can be disabled when multipliers and/or
// stacking patch cables are used. See the configuration struct passed to this
// function for more information
func (sf SynthdownFile) Validate(config *ValidationConfiguration) (err error) {
	namedJacks := make(moduleJackMapping)

	for _, patch := range sf.Patches {
		err = patch.Validate()
		if err != nil {
			return
		}

		for _, module := range patch.Modules {
			if module.Input != nil {
				namedJacks.addMapping(module.Name, module.Input.Name, module.Input.Pos)
			}

			if module.Output != nil {
				namedJacks.addMapping(module.Name, module.Output.Name, module.Output.Pos)
			}
		}
	}

	return namedJacks.errors()
}
