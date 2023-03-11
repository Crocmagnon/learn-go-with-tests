package _7_maps

import "testing"

var word = "test"
var definition = "this is just a test"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(word)
		want := definition

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add(word, definition)
		assertNoError(t, err)

		got, err := dict.Search(word)
		want := definition
		assertNoError(t, err)
		assertStrings(t, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new def")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: definition}
		newDef := "new definition"
		err := dict.Update(word, newDef)
		assertNoError(t, err)
		assertDefinition(t, dict, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(tb testing.TB, dictionary Dictionary, word, want string) {
	tb.Helper()
	got, err := dictionary.Search(word)
	assertNoError(tb, err)
	if got != want {
		tb.Errorf("got %q want %q given %q", got, want, word)
	}
}

func assertError(tb testing.TB, got, want error) {
	tb.Helper()
	if got != want {
		tb.Errorf("got error %q, want %q", got, want)
	}
}

func assertNoError(tb testing.TB, got error) {
	tb.Helper()
	assertError(tb, got, nil)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
