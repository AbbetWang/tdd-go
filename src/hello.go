package main

import "fmt"

func main() {
	x := Hello("world", "")
	fmt.Println(x)
}

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frechHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frechHelloPrefix + name
	}
	return englishHelloPrefix + name
}
