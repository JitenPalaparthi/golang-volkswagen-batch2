package main

import "fmt"

func main() {

	var r1 Rect
	r1.L = 12.23
	r1.B = 12.23

	a1 := Area(r1)
	p1 := Perimeter(r1)

	fmt.Println("Area of Rect r1:", a1)
	fmt.Println("perimeter of Rect r1:", p1)

	a2 := r1.Area()
	//(&r1).Area()
	p2 := r1.Perimeter()

	fmt.Println("Area of Rect r1:", a2)
	fmt.Println("perimeter of Rect r1:", p2)
	fmt.Println("Area:", r1.A)
	fmt.Println("Perimeter:", r1.P)

	r2 := new(Rect)
	r2.L = 12.23
	r2.B = 12.23
	a3 := r2.Area()
	p3 := r2.Perimeter()

	fmt.Println("Area of Rect r2:", a3)
	fmt.Println("perimeter of Rect r2:", p3)

	r3 := &Rect{L: 123.23, B: 344.34}
	a4 := r3.Area()
	p4 := r3.Perimeter()

	fmt.Println("Area of Rect r3:", a4)
	fmt.Println("perimeter of Rect r3:", p4)

	r4 := New(12.23, 23.23)
	a5, p5 := r4.Area(), r4.Perimeter()
	fmt.Println("Area of Rect r4:", a5)
	fmt.Println("perimeter of Rect r4:", p5)

}

type integer = int // aliasing the same type. Both of them are same in all aspects

// normal struct
type Rect struct { // This is a new type
	L    float32
	B    float32
	A, P float32
}

type Square struct {
	S float32
}

func Area(r Rect) float32 {
	return r.L * r.B
}

func Perimeter(r Rect) float32 {
	return 2 * (r.L + r.B)
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func NewSquare(s float32) *Square {
	return &Square{S: s}
}

func NewDefaults() *Rect {
	return &Rect{L: 1, B: 1}
}
func NewFrom() *Rect {
	return &Rect{L: 1, B: 1}
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}
