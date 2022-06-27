package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // await Promise.all(booksService.getBook(), libraries.getLibrary())

	go func() {
		escrever("Ola mundo")
		waitGroup.Done() // - 1
	}()
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
