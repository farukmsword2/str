//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Take returns the first length runes of the string (clamped).

	// Example: take head
	v := str.Of("gophers").Take(3).String()
	godump.Dump(v)
	// #string gop
}
