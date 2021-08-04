package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		// want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expect to get an error")

		}
		assertError(t, err, ErrorNotFound)
	})

}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
	t.Run("exists word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find the add word", err)
	}
	if got != definition {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExists)
	})
	t.Run("exists word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
		assertError(t, err, nil)
	})

}

func TestDelete(t *testing.T) {
	word := "word"
	definition := "test definition"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrorNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}
