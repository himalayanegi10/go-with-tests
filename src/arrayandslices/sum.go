package main

import (

)


func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}