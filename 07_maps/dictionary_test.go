package _7_maps

import "testing"

func TestSearch(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := map[string]string{word: definition}

	got := Search(dictionary, word)
	want := definition

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
