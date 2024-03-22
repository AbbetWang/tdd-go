package main

import "fmt"

func main() {
	x := Hello("world")
	fmt.Println(x)
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if "" == name {
		name = "World"
	}
	return englishHelloPrefix + name
}
