package main

func Sum(arrays [5]int) int {
	result := 0
	for i := range arrays {
		result += arrays[i]
	}
	return result
}
