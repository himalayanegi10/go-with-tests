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

func TestReduce(t *testing.T) {
	t.Run("Test Reduce", func(t *testing.T) {
		sum := func(a, b int) int {
			return a + b
		}
		test_array := []int{1,2,3,4,5};

		got := Reduce(test_array, sum, 0)
		want := 15

		if got != want {
			t.Errorf("got %v, expected %v", got, want)
		}
	})

	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1,2,3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"Hello", "World"}, concatenate, ""), "HelloWorld")
		
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}