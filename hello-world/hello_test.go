package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people, when name is provided", func(t *testing.T) {
		got := Hello("Suraj")
		want := "Hello, Suraj"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to the world, when name is not provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
