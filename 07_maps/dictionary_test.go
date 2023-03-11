package _7_maps

import (
	"fmt"
	"testing"
)

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

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: definition}
		dict.Delete(word)
		_, err := dict.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{}
		dict.Delete(word)
	})
}

func ExampleDictionary() {
	dict := Dictionary{}

	_, err := dict.Search("capybara")
	if err == ErrNotFound {
		fmt.Println("Couldn't find word in dict.")
	}

	_ = dict.Add("capybara", "The capybara is a giant cavy rodent native to South America.")
	fmt.Println(dict)

	err = dict.Add("capybara", "again")
	if err == ErrWordExists {
		fmt.Println("Word already exists in dict")
	}

	res, _ := dict.Search("capybara")
	fmt.Println(res)

	err = dict.Update("cavy", "")
	if err == ErrWordDoesNotExist {
		fmt.Println("Word does not exist in dict")
	}

	_ = dict.Update("capybara", "new definition")
	res, _ = dict.Search("capybara")
	fmt.Println(res)

	// Output:
	// Couldn't find word in dict.
	// map[capybara:The capybara is a giant cavy rodent native to South America.]
	// Word already exists in dict
	// The capybara is a giant cavy rodent native to South America.
	// Word does not exist in dict
	// new definition
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
