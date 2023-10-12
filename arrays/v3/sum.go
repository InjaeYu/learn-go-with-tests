package main

func Sum(arrays []int) int {
	result := 0
	for _, number := range arrays {
		result += number
	}
	return result
}
