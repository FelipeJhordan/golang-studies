package shapes

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 0
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}
