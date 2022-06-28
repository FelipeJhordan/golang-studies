package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u1 := usuario{
			"Felipe",
			"felipejordan.alves@gmail.com",
		}
		templates.ExecuteTemplate(w, "index.html", u1)
	})

	fmt.Println("Escuta na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
