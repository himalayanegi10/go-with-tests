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
		got := rectangle.Perimeter()
		want := 22.0

		if got != want {
			t.Errorf("Expected %.2f got %.2f", want, got)
		}
	})

	t.Run("Test Area of a rectangle using struct", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.0}
		got := rectangle.Area()
		want := 30.0

		if got != want {
			t.Errorf("Expected %.2f got %.2f", want, got)
		}
	})

	t.Run("Test Area of a Circle using struct", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Expected %g got %g", want, got)
		}
	})
}


func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected %g got %g", want, got)
		}
	}

	t.Run("Area of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		checkArea(t, rectangle, 200.0)
	})

	t.Run("Area of Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}