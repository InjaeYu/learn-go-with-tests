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

	return greetingPrefic(language) + word
}

func greetingPrefic(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefic
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
