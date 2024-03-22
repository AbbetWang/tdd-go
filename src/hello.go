package main

import "fmt"

func main() {
	x := hello()
	fmt.Println(x)
}

func hello() string {
	x := "Hello, world"
	return x
}
