package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {

	dictionary := Dictionary{Dict: map[string]string{"test": "this is a test string"}}

	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}

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