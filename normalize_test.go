package str

import "testing"

func TestNormalizeSpace(t *testing.T) {
	t.Parallel()

	if got := Of("  go\t forj  ").NormalizeSpace().String(); got != " go forj " {
		t.Fatalf("NormalizeSpace = %q", got)
	}
	if got := Of("   ").NormalizeSpace().String(); got != " " {
		t.Fatalf("NormalizeSpace all-space = %q", got)
	}
	if got := Of("").NormalizeSpace().String(); got != "" {
		t.Fatalf("NormalizeSpace empty = %q", got)
	}
}

func TestNormalizeNewlines(t *testing.T) {
	t.Parallel()

	in := "a\r\nb\rc\u0085d\u2028e\u2029f\n"
	want := "a\nb\nc\nd\ne\nf\n"
	if got := Of(in).NormalizeNewlines().String(); got != want {
		t.Fatalf("NormalizeNewlines = %q", got)
	}
	if got := Of("a\nb").NormalizeNewlines().String(); got != "a\nb" {
		t.Fatalf("NormalizeNewlines no change = %q", got)
	}
}

func TestLines(t *testing.T) {
	t.Parallel()

	got := Of("a\r\nb\nc").Lines()
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Fatalf("Lines = %#v", got)
	}
}
