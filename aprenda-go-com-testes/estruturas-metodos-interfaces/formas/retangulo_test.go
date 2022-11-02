package formas

import "testing"

func TestPerimetro(t *testing.T) {
	result := Perimetro(10.0, 10.0)
	expect := 40.0

	if result != expect {
		t.Errorf("resultado %.2f esperado %.2f", result, expect)
	}
}
