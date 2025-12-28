//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Of wraps a raw string with fluent helpers.

	// Example: wrap raw string
	v := str.Of("gopher")
	println(v.String())
	// #string gopher
}
