//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Join joins the provided words with sep.

	// Example: join words
	v := str.Of("unused").Join([]string{"foo", "bar"}, "-").String()
	println(v)
	// #string foo-bar
}
