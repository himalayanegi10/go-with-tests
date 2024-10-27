package main

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1,2,3,4,5,6,7,8,9,10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x % 2 == 0
		})
		AssertTrue(t, found)		
		AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("find the best ufc fighter", func(t *testing.T) {
		fighters := []Fighter{
			{Name: "Alex Pereira", aka: "Po atan"},
			{Name: "Jon Jones", aka: "Bones"},
			{Name: "Ilia Topuria", aka: "El Matador"},
		}

		p4p, found := Find(fighters, func(f Fighter) bool {
			return strings.Contains(f.Name, "Ilia")
		})

		AssertTrue(t, found)
		AssertEqual(t, p4p, Fighter{Name: "Ilia Topuria", aka: "El Matador"})
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
