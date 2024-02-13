package main

import (
	"strconv"
)

type Value struct {
	Str   *string  `  @String`
	Float *float64 `| @Float`
	Int   *int     `| @Int`
}

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
