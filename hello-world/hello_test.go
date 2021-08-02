package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Dev"
	got := Hello("Dev")

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}