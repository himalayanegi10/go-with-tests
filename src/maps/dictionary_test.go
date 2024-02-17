package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	t.Run("Search a word", func(t *testing.T){

		// dictionary := map[string]string{"test": "this is a test string"}
		// got := Search(dictionary, "test")
		dictionary := Dictionary{Dict: map[string]string{"test": "this is a test string"}}
		got := dictionary.Search("test")
		want := "this is a test string"
		assertStrings(t, got, want)
		
	})
}