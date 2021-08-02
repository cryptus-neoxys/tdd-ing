package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}

	return getLangPrefix(lang) + name
}

func getLangPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Dev", "English"))
}
