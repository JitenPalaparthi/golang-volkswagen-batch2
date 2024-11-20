package main

import (
	"fmt"
	"unsafe"
)

func main() {

	c1 := NewColourCode(100, "red", "some red")

	fmt.Println(*c1)

	var e1 Empty
	e1.Greet()

	e2 := Empty(struct{}{}) // empty struct and type casted to Empty
	e2.Greet()

	var e3 struct{} // a variable is created
	e3 = struct{}{}
	fmt.Println(e3)
	if &e3 != nil {
		println("yes")
	}
	fmt.Println("Size of ", unsafe.Sizeof(&e3))
	fmt.Printf("Address of e1:%p\n", &e1)
	fmt.Printf("Address of e2:%p\n", &e2)
	fmt.Printf("Address of e3:%p\n", &e3)

	fmt.Printf("Address of c1:%p\n", &c1)

	var num = 100
	fmt.Printf("Address of num:%p\n", &num)

	var ptr1 *struct{} = &e3

	*ptr1 = struct{}{}

	//var e5 empty1
	a1 := empty1{}.Area(10.12, 13.14)
	a2 := empty2{}.Area(123.23)
	a3 := empty3{}.Area(123.23)

	println(a1, a2, a3)

}

type mystring = string

type ColorCode struct { // struct with anonymous fileds
	int
	string
	mystring
}

func NewColourCode(i int, s string, ms mystring) *ColorCode {
	return &ColorCode{int: i, string: s, mystring: ms}
}

type Empty struct{} // empty struct

func (e *Empty) Greet() {
	println("This is just an empty struct")
}

type empty1 struct{}
type empty2 struct{}
type empty3 struct{}

func (e empty1) Area(l, b float32) float32 {
	return l * b
}

func (e empty2) Area(s float32) float32 {
	return s * s
}

func (e empty3) Area(r float32) float32 {
	return 3.14 * r * r
}
