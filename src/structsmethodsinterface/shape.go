package structsmethodsinterfaces

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}