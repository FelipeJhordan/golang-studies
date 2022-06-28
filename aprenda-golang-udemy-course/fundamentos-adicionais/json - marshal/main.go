package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Racao string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}

	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c) // transformar cachorro em json, porem vem em um slice de bit
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // transformar os bytes em um json propriamente dito

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJson)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))
	fmt.Println(c2)
}
