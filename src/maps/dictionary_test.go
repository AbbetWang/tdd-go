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
		_, got := dictionary.Search("unknown")
		want := ErrNotFound

		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, want)
	})
}
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, want)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
