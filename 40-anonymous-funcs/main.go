package main

import "fmt"

var f1 func(int, int) int = func(a, b int) int {
	return a + b
}

var s int = sub(100, 200)

func init() {
	println("sub:", s)
}

func init() {
	println("anoter init-2")
}

func init() {
	println("anoter init-3")
}

//var global int = 100

func main() {
	func() {
		println("hello world")
	}() // executor

	func(msg string) {
		println(msg)
	}("Hello VolksWagen")

	ok := func(age uint8) bool {
		if age >= 18 {
			return true
		}
		return false
	}(199)
	fmt.Println("oka:", ok)

	sum1 := func(a, b int) int {
		return a + b
	}(100, 200)
	fmt.Println("sum1:", sum1)

	//var f1 func(int, int) int

	// f1 := func(a, b int) int {
	// 	return a + b
	// }

	sum2 := f1(123, 123)
	fmt.Println("sum2:", sum2)

}

func sub(a, b int) int {
	return a - b
}
