package main

func Sum(numbers []int) (result int) {
	add := func(a, b int) int { return a + b}
	return Reduce(numbers, add)
}

func Reduce[T any](array []T, F func(a, b T) T) T {
	var result T
	if(len(array) > 0) {
		result = F(array[0], Reduce(array[1:], F))
	}
	return result
}