//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// TrimRight trims trailing characters. If cutset is the zero value (empty string), trims Unicode whitespace.

	// Example: trim right
	v := str.Of("  GoForj  ").TrimRight("").String()
	println(v)
	// #string   GoForj
}
