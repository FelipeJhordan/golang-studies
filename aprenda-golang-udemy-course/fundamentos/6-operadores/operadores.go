package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// FIM ARITIMETICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// Fim operatores atribuição

	// Operadores relacionais são iguais ao javascript

	// Operadores logicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
	// Fim operadores logicos

	// Operador ternário

}
