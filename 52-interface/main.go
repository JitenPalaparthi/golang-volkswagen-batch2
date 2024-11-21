package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a1 Empty
	if a1 == nil {
		println("yes nil")
	}
	// v := a1.(int)
	fmt.Println("size of a1:", unsafe.Sizeof(a1))
	a1 = 100
	fmt.Println(a1)
	a1 = "hello world"
	fmt.Println(a1)
	a1 = true
	fmt.Println(a1)
}

type Empty interface { // can also create empty interface that can satisfy anything
}
