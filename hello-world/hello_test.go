package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Suraj")
	want := "Hello, Suraj"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
