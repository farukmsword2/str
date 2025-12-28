//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Dump is a thin alias for godump.Dump, provided to keep examples concise
	// and readable in documentation and code snippets.

	// Example: debug print
	str.Dump("go", 42)
	// #string go
	// #int 42
}
