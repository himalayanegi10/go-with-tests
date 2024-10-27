package main

func Sum(numbers []int) (result int) {
	add := func(a, b int) int { return a + b}
	return Reduce(numbers, add, 0)
}

func Reduce[T any](array []T, F func(a, b T) T, accumulator T) T {
	var result = accumulator
	for _, x := range array {
		result = F(result, x)
	}
	return result
}