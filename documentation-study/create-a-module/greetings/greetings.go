package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) // Aloca memória para uma determinada estrutura de dados
	for _, name := range names {        // ignora o parametro de index
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func Hello(name string) (string, error) {
	const EMPTY_STRING = ""
	if name == EMPTY_STRING {
		return name, errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v, Welcome! ", name) // Declare and inicializate

	var message2 string
	message2 = fmt.Sprintf(randomFormat(name)) // It´s same thing that above code

	returnedMessage := message + message2

	return returnedMessage, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat(name string) string {
	formats := []string{
		"How are you.",
		fmt.Sprintf("Great to see you, %s", name), // I think that golang should have a template literal like es6
		"You are ok?",
	}

	return formats[rand.Intn(len(formats))]
}
