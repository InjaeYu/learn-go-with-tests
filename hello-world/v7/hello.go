package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefic = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func Hello(word, language string) string {
	if len(word) == 0 {
		word = "World"
	}

	prefix := englishHelloPrefic

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + word
}

func main() {
	fmt.Println(Hello("world", ""))
}
