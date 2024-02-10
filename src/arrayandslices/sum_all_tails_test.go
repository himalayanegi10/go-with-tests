package main

import (
	"testing"
	"reflect"
)

func TestSumAllTrails(t *testing.T){
	t.Run("Get sum of all tail elements", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2, 3}, []int {0, 9})
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v got %v", want, got)
		}
	})

	t.Run("Element with slice size 0", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int {0, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}