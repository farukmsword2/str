package str

import "testing"

func TestEndsWith(t *testing.T) {
	t.Parallel()

	val := Of("gopher")
	if !val.EndsWith("her") || val.EndsWith("cat") {
		t.Fatalf("EndsWith mismatch")
	}
	if val.EndsWith() {
		t.Fatalf("EndsWith empty slice should be false")
	}
	if !val.EndsWithFold("HER") {
		t.Fatalf("EndsWithFold mismatch")
	}
	if val.EndsWithFold() {
		t.Fatalf("EndsWithFold empty slice should be false")
	}
	if val.EndsWithFold("CAT") {
		t.Fatalf("EndsWithFold non-match should be false")
	}
}
