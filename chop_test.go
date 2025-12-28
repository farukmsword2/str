package str

import "testing"

func TestChopStartEnd(t *testing.T) {
	t.Parallel()

	if got := Of("https://goforj.dev").ChopStart("https://", "http://").String(); got != "goforj.dev" {
		t.Fatalf("ChopStart = %q", got)
	}
	if got := Of("file.txt").ChopEnd(".txt", ".md").String(); got != "file" {
		t.Fatalf("ChopEnd = %q", got)
	}
	if got := Of("gopher").ChopStart("").String(); got != "gopher" {
		t.Fatalf("ChopStart empty prefix")
	}
	if got := Of("gopher").ChopEnd(".zip").String(); got != "gopher" {
		t.Fatalf("ChopEnd no match")
	}
	if got := Of("note.md").ChopEnd(".txt", ".md").String(); got != "note" {
		t.Fatalf("ChopEnd second suffix")
	}
	if got := Of("gopher").ChopEnd().String(); got != "gopher" {
		t.Fatalf("ChopEnd empty list")
	}
	if got := Of("gopher").ChopEnd("").String(); got != "gopher" {
		t.Fatalf("ChopEnd empty suffix skip")
	}
}
