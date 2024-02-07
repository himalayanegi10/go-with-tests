package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing Sum", func(t *testing.T){
		numbers := [5]int{1,2,3,4,5}
		sum := Sum(numbers)
		want := 15
		if sum != want {
			t.Errorf("Expected %d, found %d, numbers array=%v", want, sum, numbers)
		}
		// assertEquals(t, sum, want, numbers)
	}) 
}

// func assertEquals(t *testing.T, sum int, want int, numbers [...]int) {
// 	t.Helper()
// 	if sum != want {
// 		t.Errorf("Expected %d, found %d, numbers array=%v", want, sum, numbers)
// 	}
// }