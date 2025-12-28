package str

import "github.com/goforj/godump"

// Dump is a thin alias for godump.Dump, provided to keep examples concise
// and readable in documentation and code snippets.
// @group Debug
//
// Example: debug print
//
//	str.Dump("go", 42)
//	// #string go
//	// #int 42
func Dump(vs ...any) {
	godump.Dump(vs...)
}
