package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}
func main() {
	fmt.Println(Hello("", "Spanish"))
	fmt.Println(Hello("Elodie", "Spanish"))
	fmt.Println(Hello("", "French"))
	fmt.Println(Hello("Elodie", "French"))
}
