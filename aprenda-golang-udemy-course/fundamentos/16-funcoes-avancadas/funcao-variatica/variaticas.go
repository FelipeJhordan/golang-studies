package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {

		fmt.Println(numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 100)
	escrever("Olá mundo", 10, 10, 321321)
	fmt.Println(totalDaSoma)
}
