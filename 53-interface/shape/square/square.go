package square

type Square float32

func (r Square) Area() float32 {
	return float32(r * r)
}

func (r Square) Perimeter() float32 {
	return 4.0 * float32(r)
}

func (r Square) What() string {
	return "Square"
}
