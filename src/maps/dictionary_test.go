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

	t.Run("Add an key:value pair", func(t *testing.T) {
		dictionary.Add("name", "Himalaya Singh Negi")

		got, err := dictionary.Search("test")
		want := "this is a test string"

		if err != nil {
			t.Fatal("should find added word:", err)
		}
		// assertError(t, err, ErrorNotFound)
		assertStrings(t, got, want)
	})
}