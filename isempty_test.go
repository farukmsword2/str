package str

import "testing"

func TestIsEmptyBlank(t *testing.T) {
	t.Parallel()

	if !Of("").IsEmpty() {
		t.Fatalf("IsEmpty expected true")
	}
	if !Of(" \t\n").IsBlank() {
		t.Fatalf("IsBlank expected true")
	}
	if Of("go").IsBlank() {
		t.Fatalf("IsBlank expected false")
	}
}
