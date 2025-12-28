package str

import "testing"

func TestSnake(t *testing.T) {
	t.Parallel()

	if got := Of("fooBar baz").Snake("_").String(); got != "foo_bar_baz" {
		t.Fatalf("Snake = %q", got)
	}
	if got := Of("fooBar").Snake("").String(); got != "foo_bar" {
		t.Fatalf("Snake default sep = %q", got)
	}
}
