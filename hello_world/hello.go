package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	prefixMap := map[string]string{
		french:  frenchHelloPrefix,
		spanish: spanishHelloPrefix,
	}
	prefix, present := prefixMap[language]
	if !present {
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Gab", "French"))
}
