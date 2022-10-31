package iteracao

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	quantity := int64(4)
	repititions := Repeat("a", quantity)
	expected := "aaaa"

	if repititions != expected {
		t.Errorf("esperado %s mas obteve '%s'", expected, repititions)
	}
}

func ExampleRepeat() {
	quantity := int64(4)
	repititions := Repeat("a", quantity)

	fmt.Print(repititions)
	//  Output: aaaa
}
