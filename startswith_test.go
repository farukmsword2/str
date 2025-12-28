package str

import "testing"

func TestStartsWith(t *testing.T) {
	t.Parallel()

	val := Of("gopher")
	if !val.StartsWith("go") || val.StartsWith("rust") {
		t.Fatalf("StartsWith mismatch")
	}
	if val.StartsWith() {
		t.Fatalf("StartsWith empty slice should be false")
	}
	if !val.StartsWithFold("GO") {
		t.Fatalf("StartsWithFold mismatch")
	}
	if val.StartsWithFold() {
		t.Fatalf("StartsWithFold empty slice should be false")
	}
	if val.StartsWithFold("CAT") {
		t.Fatalf("StartsWithFold non-match should be false")
	}
}
