package main

import (
	"os"
	"time"

	"example.com/tdd-go/src/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
