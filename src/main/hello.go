package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const exclamation = " !"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "spanish" {
		return spanishHelloPrefix + name + exclamation
	}

	if language == "french" {
		return frenchHelloPrefix + name + exclamation
	}

	return englishHelloPrefix + name + exclamation
}

func main() {
	fmt.Println(Hello("world", ""))
}