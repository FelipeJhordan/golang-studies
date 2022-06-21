package main

import (
	"fmt"
	"modulos-externos/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.EscreverAlgumaCoisa()

	erro := checkmail.ValidateFormat("felipejordan.alves@gmail.com")
	fmt.Println(erro)
}
