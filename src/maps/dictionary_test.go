package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test", "a": "this is just a a"}
	t.Run("happy path", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, "this is just a test")
		got, _ = dictionary.Search("a")
		assertStrings(t, got, "this is just a a")
	})
	t.Run("sad path,if a key given is not in the map", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
