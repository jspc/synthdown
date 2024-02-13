package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
)

type errFirstPatchHasInput struct {
	m Module
}

func (e errFirstPatchHasInput) Error() string {
	return fmt.Sprintf("patch described at %s has an input", e.m.Pos)
}

type Patch struct {
	Pos lexer.Position

	Modules []Module `parser:"@@ ('-' '>' @@)*"`
}

func (p Patch) Validate() (err error) {
	modulesLen := len(p.Modules)

	if modulesLen == 0 {
		return
	}

	// Ensure first module _doesn't_ have an input
	first := p.Modules[0]
	if first.Input != nil {
		return errFirstPatchHasInput{first}
	}

	return
}
