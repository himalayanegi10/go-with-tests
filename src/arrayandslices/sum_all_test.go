package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("Single Slice as input", func(t *testing.T) {
		got := SumAll([]int{1,1,1}, []int{5,6})
		want := []int{3, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})
}