package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 1000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var texto string
	fmt.Println(texto)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
