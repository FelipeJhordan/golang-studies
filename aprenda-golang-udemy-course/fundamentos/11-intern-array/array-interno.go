package main

import "fmt"

func main() {
	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	fmt.Println((len(slice))) // length
	fmt.Println((cap(slice))) // capacidade

	slice = append(slice, 213)
	fmt.Println(slice)
	slice = append(slice, 213)
	fmt.Println(slice)

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
}
