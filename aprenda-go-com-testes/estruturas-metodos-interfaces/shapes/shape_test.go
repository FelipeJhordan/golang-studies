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
	verifyArea := func(t *testing.T, shape Shape, expect float64) {
		t.Helper()
		result := shape.Area()

		if result != expect {
			t.Errorf("esperado %.2f, esperado %.2f", result, expect)
		}
	}

	t.Run("retangulos", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}

		expect := 72.00

		verifyArea(t, rectangle, expect)

	})

	t.Run("círculos", func(t *testing.T) {
		circle := Circle{10}

		expect := 314.1592653589793

		verifyArea(t, circle, expect)

	})
}
