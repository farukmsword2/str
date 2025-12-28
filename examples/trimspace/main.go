//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// TrimSpace trims leading and trailing Unicode whitespace.

	// Example: trim space
	v := str.Of("  GoForj  ").TrimSpace().String()
	println(v)
	// #string GoForj
}
