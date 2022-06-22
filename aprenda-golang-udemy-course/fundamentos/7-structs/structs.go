package main

import "fmt"

type Usuario struct {
	nome     string
	idade    uint8
	endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u Usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	usuario2 := Usuario{
		"Felipe",
		21,
		Endereco{
			logradouro: "Rua dos bobos",
			numero:     12,
		},
	}
	fmt.Println(usuario2)

	usuario3 := Usuario{
		idade: 21,
	}

	fmt.Println(usuario3)
}
