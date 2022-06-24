package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 0:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 1:
		diaDaSemana = "Segunda-feira"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}

func main() {
	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
