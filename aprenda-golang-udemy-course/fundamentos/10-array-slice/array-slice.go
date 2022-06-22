package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 0"
	fmt.Println(array1)

	array := [5]string{"Posição1", "Posição2", "Posição3", "Posição 4", "Posição 5"}
	fmt.Println(array)
	// array[5] = "Posição 6"

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := array3[1:2]
	fmt.Println(slice2)
}
