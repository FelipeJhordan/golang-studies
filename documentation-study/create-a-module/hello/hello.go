package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"Felipe", "Raquel", "Samuel", "Christian"}

	message, err := greetings.Hellos(names)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
