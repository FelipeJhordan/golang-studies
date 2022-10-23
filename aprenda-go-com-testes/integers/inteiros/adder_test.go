package inteiros

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sumResult := Adder(2, 2)
	expected := 4

	if sumResult != expected {
		t.Errorf("esperado '%d', resultado '%d'", expected, sumResult)
	}
}


func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Print(sum)
	// Output: 6
}
