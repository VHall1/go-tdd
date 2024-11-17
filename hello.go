package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	// spanish
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	// french
	french            = "French"
	franchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = franchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
