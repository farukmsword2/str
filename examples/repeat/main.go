//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Repeat repeats the string count times (non-negative).

	// Example: repeat string
	v := str.Of("go").Repeat(3).String()
	godump.Dump(v)
	// #string gogogo
}
