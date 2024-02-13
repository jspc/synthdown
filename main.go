package main

import (
	"fmt"
)

var (
	//in = `noise(white) -> (in)rand[L:10, R:10](out) -> (in)vca[L:10](out)`
	in = `noise[](white) -> (in)rand[L:10, R:10](out) -> (in)vcs[L:10](out);`
)

func main() {
	_, err := New(in)
	fmt.Println(err)
}
