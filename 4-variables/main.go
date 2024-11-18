package main

import "fmt"

var Global int = 100 // data segment

func main() {
	const PI float32 = 3.14 // code segment
	const MAX = 7           // inference and kind of a shorthand declaration for constants
	const MIN int = 0       // there is no zero value for constant, should assign a value to a constant

	var num = 1
	// group of constants
	const (
		A = 9999
		B = A + 100 //+ num
		C = (A + B) * 2
		D = 100*100 + ((A+B)/100)%2
	)
	str1 := "Hello World, how are you doi'n" // does it store in stack or heap
	fmt.Println(str1, num)
	Global++ // This can be changed bcz it is a global variable stored in data segment
	// D = 300 // cant be changed bcz it is a constant
}

// can I give an expression to a constant?
