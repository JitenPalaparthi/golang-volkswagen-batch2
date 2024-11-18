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
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))
	any1 = "Hello World!"
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))
	any1 = true
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))
	any1 = 12312.12312
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	if any1 == nil {
		fmt.Println("any1 is nil")
	} else {
		fmt.Println("any1 is not nil")
	}
}

/* any1
ptr: 8 bytes 0x122122
type:8 bytes float64
*/
