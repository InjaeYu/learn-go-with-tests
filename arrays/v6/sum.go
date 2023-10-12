package main

func Sum(arrays []int) int {
	result := 0
	for _, number := range arrays {
		result += number
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers[1:]))
	}

	return sums
}
