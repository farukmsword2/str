//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ReplaceLastFold replaces the last occurrence of old with repl using Unicode case folding.

	// Example: replace last (case-insensitive)
	v := str.Of("go gopher GO").ReplaceLastFold("GO", "Go").String()
	godump.Dump(v)
	// #string go gopher Go
}
