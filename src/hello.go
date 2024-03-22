package main

import "fmt"

func main() {
	x := Hello()
	fmt.Println(x)
}

func Hello() string {
	x := "Hello, world"
	return x
}
