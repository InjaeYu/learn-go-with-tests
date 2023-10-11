package main

import "fmt"

const englishHelloPrefic = "Hello, "

func Hello(word string) string {
	if len(word) == 0 {
		word = "World"
	}
	return englishHelloPrefic + word
}

func main() {
	fmt.Println(Hello("world"))
}
