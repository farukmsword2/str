package str

import "testing"

func TestReverse(t *testing.T) {
	t.Parallel()

	if got := Of("naïve").Reverse().String(); got != "evïan" {
		t.Fatalf("Reverse = %q", got)
	}
}
