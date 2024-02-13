package synthdown

import (
	"strconv"
)

// Value represents, typically, the value part of an arg. Or,
// for example, the `3` in the args `[A: 3]`.
//
// Because an argument can typically be an integer, float, or a
// string we must support each type.
//
// The zero value (ie: when none of the inner values are set) is a 0
type Value struct {
	Str   *string  `  @String`
	Float *float64 `| @Float`
	Int   *int     `| @Int`
}

// Float64 returns the float64 representation of a Value, casting appropriately
// so that:
//  1. A float64 is returned as is
//  2. An int is cast as a float64; and
//  3. A string is `strconv`d as a float64
//
// This function panics if a string is passed in which cannot be `strconv`d as a float64.
func (v Value) Float64() float64 {
	if v.Float != nil {
		return *v.Float
	}

	if v.Int != nil {
		return float64(*v.Int)
	}

	if v.Str == nil {
		return 0
	}

	out, err := strconv.ParseFloat(*v.Str, 64)
	if err != nil {
		panic(err)
	}

	return out
}

// String returns a string representation of the value, which is especially
// useful for diagramming where you don't really care about what's set
func (v Value) String() string {
	if v.Str != nil {
		return *v.Str
	}

	if v.Float != nil {
		return strconv.FormatFloat(*v.Float, 'f', -1, 64)
	}

	if v.Int == nil {
		return "0"
	}

	return strconv.Itoa(*v.Int)
}
