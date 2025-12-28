package str

import "testing"

func TestEquals(t *testing.T) {
	t.Parallel()

	val := Of("gopher")
	if !val.Equals("gopher") || val.Equals("Gopher") {
		t.Fatalf("Equals mismatch")
	}
	if !val.EqualsFold("GOPHER") {
		t.Fatalf("EqualsFold mismatch")
	}
}
