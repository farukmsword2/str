//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// TakeLast returns the last length runes of the string (clamped).

	// Example: take tail
	v := str.Of("gophers").TakeLast(4).String()
	godump.Dump(v)
	// #string hers
}
