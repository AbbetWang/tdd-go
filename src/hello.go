package main

import "fmt"

func main() {
	x := Hello("world", "")
	fmt.Println(x)
}

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}
