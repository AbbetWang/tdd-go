package helloworld

import "fmt"

func main() {
	x := Hello("world", "")
	fmt.Println(x)
}

const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frechHelloPrefix   = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return grettingPrefix(language) + name
}

func grettingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frechHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
