package main

import "fmt"

func main() {
	x := Hello("world")
	fmt.Println(x)
}

func Hello(name string) string {
	x := "Hello, " + name
	return x
}
