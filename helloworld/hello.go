package helloworld


const (
	french = "french"
	spanish = "spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

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


func Hello(name string, language string ) string {

	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}


