package helloworld

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

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = franchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
