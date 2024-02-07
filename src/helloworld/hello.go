package main

import "fmt"

const (
	spanish = "spanish"
	french = "french"
	sanskrit = "sanskrit"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	sanskritHelloPrefix = "Namah te, "

	exclamation = " !"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + exclamation
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case sanskrit :
		prefix = sanskritHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}