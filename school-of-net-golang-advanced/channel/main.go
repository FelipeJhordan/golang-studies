package main

import (
	"fmt"
	"sync"
)

// vai ser executado qnd aplicação for iniciada

var waitGroup sync.WaitGroup

func main() {
	msg := make(chan string)

	go func() {
		msg <- "Hello World"
	}()

	result := <-msg
	fmt.Println(result)

}
