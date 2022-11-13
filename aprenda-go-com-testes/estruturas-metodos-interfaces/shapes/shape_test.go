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
		name   string
		shape  Shape
		expect float64
	}{
		{
			name:   "Retangulo",
			shape:  Rectangle{width: 12, height: 6},
			expect: 72.00,
		},
		{
			name: "Circle",
			shape: Circle{
				radius: 10,
			},
			expect: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				base:   12,
				height: 6,
			},
			expect: 36.00,
		},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expect {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.shape, result, tt.expect)
			}
		})

	}
}
