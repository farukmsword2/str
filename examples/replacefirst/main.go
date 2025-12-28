//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// ReplaceFirst replaces the first occurrence of old with repl.

	// Example: replace first
	v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
	println(v)
	// #string go gopher
}
