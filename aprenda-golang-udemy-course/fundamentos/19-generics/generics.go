package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(true)
	generica(1)
	generica(float32(32.2))

	mapa := map[interface{}]interface{}{
		true:  true,
		"lre": "dsa",
	}

}
