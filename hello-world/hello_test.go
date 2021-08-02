package main

import "testing"

func TestHello(t *testing.T) {
	
	assertCorrectMessage := func (t testing.TB, got, want string) {
		t.Helper()
		
		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Dev"
		got := Hello("Dev", "English")

		assertCorrectMessage(t, got, want)
	})

	t.Run(`Say "Hello, world" when empty string is passed`, func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "English")

		assertCorrectMessage(t, got, want)
	})
	
    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", spanish)
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
	
    t.Run("in French", func(t *testing.T) {
        got := Hello("Mademoiselle", french)
        want := "Bonjour, Mademoiselle"
        assertCorrectMessage(t, got, want)
    })
}