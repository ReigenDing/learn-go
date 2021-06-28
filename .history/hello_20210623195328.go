package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spnishPrefix = "Spanish, "
const frenchPrefix = "French, "

func Hello(name string, language string) string {
	// if name == "" {
	// 	name = "World"
	// }
	// if language == "Spanish" {
	// 	return "Hola, " + name
	// }
	// if language == "French" {
	// 	return "Bonjour, " + name
	// }
	// return englishHelloPrefix + name
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spnishPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
