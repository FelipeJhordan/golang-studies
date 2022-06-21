package main

import "fmt"

func main() {
	var variavel1 string = ""
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante"

	fmt.Println(constante1)

	// Inverter valores

	variavel5, variavel6 = variavel6, variavel5
}
