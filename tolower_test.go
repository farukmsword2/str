package str

import "testing"

func TestToLower(t *testing.T) {
	t.Parallel()

	if got := Of("GoLang").ToLower().String(); got != "golang" {
		t.Fatalf("ToLower = %q", got)
	}
}
