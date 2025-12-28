package str

import "testing"

func TestToTitle(t *testing.T) {
	t.Parallel()

	if got := Of("ß").ToTitle().String(); got != "SS" {
		t.Fatalf("ToTitle = %q", got)
	}
	if got := Of("").ToTitle().String(); got != "" {
		t.Fatalf("ToTitle empty = %q", got)
	}
	if got := Of("go").ToTitle().String(); got != "GO" {
		t.Fatalf("ToTitle letters = %q", got)
	}
}
