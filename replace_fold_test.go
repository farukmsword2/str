package str

import "testing"

func TestReplaceFoldVariants(t *testing.T) {
	t.Parallel()

	val := Of("go gopher GO")
	if got := val.ReplaceFold("GO", "Go").String(); got != "Go Gopher Go" {
		t.Fatalf("ReplaceFold = %q", got)
	}
	if got := val.ReplaceFirstFold("GO", "Go").String(); got != "Go gopher GO" {
		t.Fatalf("ReplaceFirstFold = %q", got)
	}
	if got := val.ReplaceLastFold("GO", "Go").String(); got != "go gopher Go" {
		t.Fatalf("ReplaceLastFold = %q", got)
	}
	if got := val.ReplaceFold("", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceFold empty old = %q", got)
	}
	if got := val.ReplaceFirstFold("", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceFirstFold empty old = %q", got)
	}
	if got := val.ReplaceLastFold("", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceLastFold empty old = %q", got)
	}
	if got := val.ReplaceFold("missing", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceFold missing = %q", got)
	}
	if got := val.ReplaceFold("longer value", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceFold longer = %q", got)
	}
	if got := val.ReplaceFirstFold("longer", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceFirstFold longer = %q", got)
	}
	if got := val.ReplaceLastFold("nope", "Go").String(); got != val.String() {
		t.Fatalf("ReplaceLastFold missing = %q", got)
	}
}

func TestReplaceFoldInternals(t *testing.T) {
	t.Parallel()

	if got, ok := replaceFoldAll("go", "", "x"); ok || got != "go" {
		t.Fatalf("replaceFoldAll empty old = %q, %v", got, ok)
	}
	if _, _, ok := findFoldMatch("go", "", false); ok {
		t.Fatalf("findFoldMatch empty old expected false")
	}
}
