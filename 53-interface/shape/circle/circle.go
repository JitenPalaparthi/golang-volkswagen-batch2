package circle

type Circle struct {
	R float32
}

func New(r float32) Circle {
	return Circle{R: r}
}

func (c *Circle) Area() float32 {
	return 3.14 * c.R * c.R
}

func (c *Circle) Perimeter() float32 {
	return 2 * 3.14 * c.R
}

func (c *Circle) What() string {
	return "Circle"
}
