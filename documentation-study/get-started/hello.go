package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

// go run tidy procura por pacotes que estão sendo referenciados em algum arquivo.go e instala a mais recente versão para você
