package str

import "testing"

func TestSquish(t *testing.T) {
	t.Parallel()

	if got := Of("   go   forj  ").Squish().String(); got != "go forj" {
		t.Fatalf("Squish = %q", got)
	}
}
