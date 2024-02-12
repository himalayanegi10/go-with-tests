package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	// t.Run("Test Perimeter of a rectangle", func(t *testing.T) {
	// 	got := Perimeter(10.0, 20.0)
	// 	want := 60.0

	// 	if got != want {
	// 		t.Errorf("Expected %.2f got %.2f", want, got)
	// 	}
	// })

	// t.Run("Test Area of a rectangle", func(t *testing.T) {
	// 	got := Area(5.0, 6.0)
	// 	want := 30.0

	// 	if got != want {
	// 		t.Errorf("Expected %.2f got %.2f", want, got)
	// 	}
	// })

	t.Run("Test Perimeter of a rectangle using struct", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.0}
		got := Perimeter(rectangle)
		want := 22.0

		if got != want {
			t.Errorf("Expected %.2f got %.2f", want, got)
		}
	})

	t.Run("Test Area of a rectangle using struct", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.0}
		got := Area(rectangle)
		want := 30.0

		if got != want {
			t.Errorf("Expected %.2f got %.2f", want, got)
		}
	})

	// t.Run("Test Area of a Circle using struct", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	got := Area(circle)
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("Expected %g got %g", want, got)
	// 	}
	// })
}