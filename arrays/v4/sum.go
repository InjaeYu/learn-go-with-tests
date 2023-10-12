package main

func Sum(arrays []int) int {
	result := 0
	for _, number := range arrays {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
