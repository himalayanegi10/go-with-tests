package main

import "fmt"

const englishHelloPrefix = "Hello, "
const exclamation = " !"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name + exclamation
}

func main() {
	fmt.Println(Hello("world"))
}