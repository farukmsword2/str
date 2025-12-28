package str

import "testing"

func TestToUpper(t *testing.T) {
	t.Parallel()

	if got := Of("GoLang").ToUpper().String(); got != "GOLANG" {
		t.Fatalf("ToUpper = %q", got)
	}
}
