package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q found %q", expected, repeated)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	result := Repeat("b")
	fmt.Println(result)
	// Output: bbbbb
}