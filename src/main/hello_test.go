package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Himalaya")
	want := "Hello, Himalaya !"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}