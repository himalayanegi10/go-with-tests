package maps

import (
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word , definition string) {
	t.Helper()
	got , err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got , definition)	
}

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{Dict: map[string]string{"test": "this is a test string"}}

	t.Run("Search a word", func(t *testing.T){

		// dictionary := map[string]string{"test": "this is a test string"}
		// got := Search(dictionary, "test")
		got, _ := dictionary.Search("test")
		want := "this is a test string"
		assertStrings(t, got, want)
		
	})

	t.Run("Search an absent word", func(t *testing.T){

		_, err := dictionary.Search("testo")

		// if err == nil {
		// 	t.Fatal("Expected to get an error.")
		// }

		// assertStrings(t, err.Error(), want)
		assertError(t, err, ErrorNotFound)
		
	})
}


func TestAdd(t *testing.T) {
	dictionary := Dictionary{Dict: map[string]string{"test": "this is a test string"}}

	t.Run("Add new key:value pair", func(t *testing.T) {
		word := "name"
		definition := "Himalaya Singh Negi"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Add existing key:value pair", func(t *testing.T) {
		word := "test"
		definition := "this is a test string"
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{Dict: map[string]string{word: definition}}
	newDefinition := "new definition"
	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)

	t.Run("existing word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{Dict : map[string]string{word: definition}}
		newDefinition := "new Definition"

		err := dictionary.Update(word, newDefinition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	
	t.Run("new word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{Dict: map[string]string{}}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T){
	word := "test"
	dictionary := Dictionary{Dict: map[string]string{word : "test definition"}}

	err := dictionary.Delete(word)
	assertError(t, err, nil)

	_, err = dictionary.Search(word)
	assertError(t, err, ErrorNotFound)
}