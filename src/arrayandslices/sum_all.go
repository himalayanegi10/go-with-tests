package main

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, number := range numbersToSum {
	//  	sums[i] = Sum(number)
	// }

	// *-----------------------refactoring above code-----------------------*

	var sums []int
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}