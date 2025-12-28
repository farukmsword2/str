//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// UcSplit splits the string on uppercase boundaries and delimiters, returning segments.

	// Example: split on upper-case boundaries
	v := str.Of("HTTPRequestID").UcSplit()
	println(v)
	// #[]string [HTTP Request ID]
}
