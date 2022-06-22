package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Felipe",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Felipe",
			"ultimo":   "Alves",
		},
		"curso": {
			"udemy": "técnico",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "curso")

	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}

	fmt.Println(usuario2)
}
