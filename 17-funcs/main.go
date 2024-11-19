package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var F1 func()
	var F2 func(string)
	var F3 func(int, int) int

	if F1 == nil {
		println("yes F1 is nil")
	}
	F1 = Greet1
	if F1 == nil {
		println("yes F1 is nil")
	} else {
		println("No F1 is not nil, Function is assigned")
	}

	//Greet1()
	F1()
	F2 = Greet2
	F3 = Add

	F2("Hello VW Group")
	c := F3(10, 20)
	fmt.Println("Size of Greet1:", unsafe.Sizeof(Greet1))
	fmt.Println("Size of Greet2:", unsafe.Sizeof(Greet2))
	fmt.Println("C:", c, "Size of Add:", unsafe.Sizeof(Add))

	var F4 func(int, int) (int, int, int, int, int)
	if F4 == nil {
		fmt.Println("yes")
	}
	F4 = Calc
	//if Calc != nil {

	a, s, m, d, r := F4(200, 300)
	fmt.Println("Add", a, "Sub:", s, "Mul", m, "Div", d, "Rem", r)
	//}

}

func Greet1() {
	println("Hello World")
}

func Greet2(str string) {
	println(str)
}

func Add(a int, b int) int {
	return a + b
}

func Calc(a, b int) (add int, sub int, mul int, div int, rem int) {
	add, sub, mul, div, rem = a+b, a-b, a*b, a/b, a%b
	return add, sub, mul, div, rem
}
