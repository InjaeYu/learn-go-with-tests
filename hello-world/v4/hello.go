package main

import "fmt"

const englishHelloPrefic = "Hello, "

func Hello(word string) string {
	return englishHelloPrefic + word
}

func main() {
	fmt.Println(Hello("world"))
}
