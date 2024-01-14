package main

import "testing"

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bradle")
		want := "Hello, Bradley"
		assert(t, got, want)
	})

	t.Run("call Hello w/no args", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assert(t, got, want)
	})
}
