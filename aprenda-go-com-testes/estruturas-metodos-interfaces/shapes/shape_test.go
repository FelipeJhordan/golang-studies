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
	testsArea := []struct {
		shape  Shape
		expect float64
	}{
		{
			Rectangle{12, 6}, 72.00,
		},
		{
			Circle{
				10,
			}, 314.1592653589793,
		},
		{
			Triangle{
				12,
				6,
			}, 36.00,
		},
	}

	for _, tt := range testsArea {
		result := tt.shape.Area()

		if result != tt.expect {
			t.Errorf("resultado %.2f, esperado %.2f", result, tt.expect)
		}
	}
}
