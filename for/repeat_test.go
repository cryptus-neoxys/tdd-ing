package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}