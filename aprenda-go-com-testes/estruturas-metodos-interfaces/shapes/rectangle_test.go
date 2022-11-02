package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expect := 40.0

	if result != expect {
		t.Errorf("resultado %.2f esperado %.2f", result, expect)
	}
}

func TestArea(t *testing.T) {
	t.Run("retangulos", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		result := Area(rectangle)

		expect := 72.00

		if expect != result {
			t.Errorf("resultado %.2f, esperado %.2f", result, expect)
		}
	})

	t.Run("círculos", func(t *testing.T) {
		circle := Circle{10}
		result := Area(circle)

		expect := 314.1592653589793

		if result != expect {
			t.Errorf("result %.2f, expect %.2f", result, expect)
		}
	})
}
