package main

import "fmt"


const (
	french = "french"
	spanish = "spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language string ) string {

	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case french:
			prefix = frenchHelloPrefix
			break
		case spanish:
			prefix = spanishHelloPrefix
			break
		default:
			prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}