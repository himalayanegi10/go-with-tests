package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Himalaya", "")
		want := "Hello, Himalaya !"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello world if empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world !"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Himalaya", "spanish")
		want := "Hola, Himalaya !"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Himalaya", "french")
		want := "Bonjour, Himalaya !"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Sanskrit", func(t *testing.T) {
		got := Hello("Himalaya", "sanskrit")
		want := "Namah te, Himalaya !"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}