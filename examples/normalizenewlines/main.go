//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// NormalizeNewlines replaces CRLF, CR, and Unicode separators with \n.

	// Example: normalize newline variants
	v := str.Of("a\\r\\nb\\u2028c").NormalizeNewlines().String()
	godump.Dump(v)
	// #string a\nb\nc
}
