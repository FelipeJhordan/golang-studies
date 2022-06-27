package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Não irá bloquear

	canal <- "Olá mundo"
	canal <- "Olá mundo2"

	mensagem1 := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}
