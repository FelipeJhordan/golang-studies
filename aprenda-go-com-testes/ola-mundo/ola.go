package main

import "fmt"

// É importante separar minha regra de negócio dos efeitos colaterais para chegar no objetivo

const helloPrefixInPortuguese = "Olá"
const helloPrefixInEspanol = "Hola"
const helloPrefixInFrances = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "Mundo"
	}

	prefix := helloPrefixInPortuguese

	switch language {
	case "frances":
		prefix = helloPrefixInFrances
	case "espanhol":
		prefix = helloPrefixInEspanol
	}

	return prefix + ", " + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "frances":
		prefix = helloPrefixInFrances
	case "espanhol":
		prefix = helloPrefixInEspanol
	default:
		prefix = helloPrefixInPortuguese
	}

	return
}

func main() {
	fmt.Println(Hello("Felipe", ""))
}
