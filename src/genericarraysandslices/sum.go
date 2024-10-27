package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func Reduce[T comparable](array []T, F func(a, b T) T) T {
	length := len(array)
	switch(length){
	case 0:
		var zero T
		return zero
	default:
		result := F(array[0], Reduce(array[1:], F))
		return result
	}
}