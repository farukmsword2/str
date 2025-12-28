package str

import "testing"

func TestKebab(t *testing.T) {
	t.Parallel()

	if got := Of("fooBar baz").Kebab().String(); got != "foo-bar-baz" {
		t.Fatalf("Kebab = %q", got)
	}
}
