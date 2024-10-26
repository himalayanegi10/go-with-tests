package main

import (
	"testing"
	"reflect"
)

func TestSumAllTrails(t *testing.T){

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}

	t.Run("Get sum of all tail elements", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2, 3}, []int {0, 9})
		want := []int{5, 9}

		checkSums(t, got, want)
	})

	t.Run("Element with slice size 0", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int {0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}