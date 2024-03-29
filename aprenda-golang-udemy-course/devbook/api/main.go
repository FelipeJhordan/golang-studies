package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvs()

	r := router.GenerateRouter()
	fmt.Println(config.SecretKey)
	fmt.Printf("Escutando na porta %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
