package main

func walk(x interface{}, fn func(input string)) {
	fn("Chris")
}
