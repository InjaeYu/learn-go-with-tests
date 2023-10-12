package main

func Sum(arrays [5]int) int {
	result := 0
	for _, number := range arrays {
		result += number
	}
	return result
}
