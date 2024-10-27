package main

func Sum(numbers []int) (result int) {
	add := func(a, b int) int { return a + b}
	return Reduce(numbers, add, 0)
}

func Reduce[T, A any](array []T, F func(a A, b T) A, accumulator A) A {
	var result = accumulator
	for _, x := range array {
		result = F(result, x)
	}
	return result
}