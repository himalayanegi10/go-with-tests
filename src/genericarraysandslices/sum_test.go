package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing Sum", func(t *testing.T){
		numbers := []int{1,2,3,4,5}
		sum := Sum(numbers)
		want := 15
		if sum != want {
			t.Errorf("Expected %d, found %d, numbers array=%v", want, sum, numbers)
		}
		// assertEquals(t, sum, want, numbers)
	}) 

	t.Run("Testing sum by passing slice", func(t *testing.T){
		mySlice := []int{1,2,9,4,5}
		sum := Sum(mySlice)
		expected := 21

		if sum != expected {
			t.Errorf("Expected %d but found %d and slice is %v", expected, sum, mySlice)
		}
	})
}

// func assertEquals(t *testing.T, sum int, want int, numbers [...]int) {
// 	t.Helper()
// 	if sum != want {
// 		t.Errorf("Expected %d, found %d, numbers array=%v", want, sum, numbers)
// 	}
// }