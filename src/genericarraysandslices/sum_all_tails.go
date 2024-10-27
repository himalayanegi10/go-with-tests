package main

func SumAllTrails(numbersToSum ...[]int) []int {
	sumTail := func(a, b []int) []int {
		if len(b) == 0 {
			return append(a, 0)
		} else {
			tail := b[1:]
			return append(a, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}