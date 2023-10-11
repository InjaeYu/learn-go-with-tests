package main

import "fmt"

func Hello(word string) string {
	return "Hello, " + word
}

func main() {
	fmt.Println(Hello("world"))
}
