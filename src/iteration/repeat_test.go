package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMatch(t, repeated, expected)
	})

	t.Run("Repeat 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectMatch(t, repeated, expected)
	})
}

func assertCorrectMatch(t *testing.T, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("Expected %q found %q", expected, repeated)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 5)
	fmt.Println(result)
	// Output: bbbbb
}