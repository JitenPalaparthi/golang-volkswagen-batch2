package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/circle"
	"shapes/shape/cuboid"
	"shapes/shape/rect"
	"shapes/shape/square"
	"unsafe"

	"math/rand"
)

func main() {
	r1 := rect.New(12.12, 14.23)
	var s1 square.Square = 123.123
	c1 := cuboid.New(12.123, 123.23, 123.23)
	cr1 := circle.New(123.23)
	shapeSlice := make([]shape.IShape, 10)
	shapeSlice[0] = r1
	shapeSlice[1] = s1
	shapeSlice[2] = c1
	shapeSlice[3] = rect.New(123.23, 23.23)
	shapeSlice[4] = square.Square(123.123)
	shapeSlice[5] = cuboid.New(1.12, 23.3, 23.4)
	shapeSlice[6] = rect.New(1.23, 3.23)
	shapeSlice[7] = square.Square(3.12)
	shapeSlice[8] = cuboid.New(12.12, 33.3, 2.4)
	shapeSlice[9] = &cr1

	cr2 := circle.Circle{R: 123.23}
	var ishape shape.IShape = &cr2

	for i := 1; i <= 10; i++ {
		n := rand.Intn(10)
		println("rand number:", n)
		Shape(shapeSlice[n])
	}

}

func Shape(ishape shape.IShape) {
	fmt.Println("----->>>Size of I Shape:", unsafe.Sizeof(ishape))
	fmt.Println("Area of :", ishape.What(), ":", ishape.Area())
	fmt.Println("Perimeter of :", ishape.What(), ":", ishape.Perimeter())
}
