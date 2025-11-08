package main

import (
	"fmt"
)

const (
	hindi  = "Hindi"
	tamil  = "Tamil"
	german = "German"

	englishHelloPrefix = "Hello, "
	hindiHelloPrefix   = "Namaste, "
	tamilHelloPrefix   = "Namaskaram, "
	germanHelloPrefix  = "Hallo, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case hindi:
		prefix = hindiHelloPrefix
	case tamil:
		prefix = tamilHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
