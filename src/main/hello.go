package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const sanskritHelloPrefix = "Namah te, "
const exclamation = " !"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	case "sanskrit":
		prefix = sanskritHelloPrefix
	}

	return prefix + name + exclamation
}

func main() {
	fmt.Println(Hello("world", ""))
}