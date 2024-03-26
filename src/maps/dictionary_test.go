package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test", "a": "this is just a a"}

	got := Search(dictionary, "test")
	assertEqual(t, got, "this is just a test")
	got = Search(dictionary, "a")
	assertEqual(t, got, "this is just a a")
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
