package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var any1 any
	//var any2 interface{}
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))
	any1 = 100
	//var num1 int = int(any1) // cant do this
	var num float32 = float32(any1.(int)) // any.(t)
	fmt.Println(num)

	any1 = "Hello World"

	var str1 string = any1.(string)
	fmt.Println(str1)

	// arthemetic operation

	var (
		num1 int     = 100
		num2 uint8   = 200
		num3 float32 = 123.2
		num4 float64 = 123.123
		any2 any     = 100
		any3 any     = 123.123
	)

	var sum any = float64(num1) + float64(num2) + float64(num3) + num4 + float64(any2.(int)) + any3.(float64)
	fmt.Printf("sum:%.3f type of sum:%T\n", sum, sum)

	str3 := fmt.Sprintf("sum:%.3f type of sum:%T", sum, sum)
	fmt.Println(str3)

	num = (10 * 20) + 123%2 + 123 - 12 + (2+2)*2
	//    200 + 4 +1 +123 -12
	//    205+111= 316
	// 200 + 8 +1 + 111
	// 320
	fmt.Println(num)
}

// ( )
// * /
// + -
//
