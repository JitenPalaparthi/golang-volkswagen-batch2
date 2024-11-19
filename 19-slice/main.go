package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

func main() {

	var slice1 []int // this is just a declaration but not instantiations
	if slice1 == nil {
		println("yes it is")
	}
	fmt.Printf("Address of slice:%p,size of the slice:%d\n", &slice1, unsafe.Sizeof(slice1))

	slice1 = make([]int, 5) // instantiated a slice

	fmt.Printf("Address of slice:%p,size of the slice:%d\n", &slice1, unsafe.Sizeof(slice1))
	fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	fmt.Println(slice1)
	sum1 := SumOf(slice1)                                         // slice of 5 int elements
	sum2 := SumOf([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}) // slice of 10 int elements
	sum3 := SumOf([]int{})                                        // slice of 0 elements
	println(sum1)
	println(sum2)
	println(sum3)
}

func SumOf(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}
	return sum
}

// slice
// slice is a separate datastructre
// slice can grow or can shrink
// slice is a reference type, that does not mean that slice is stored in heap
// heap or stack is decided by escape analysis, not by developer
// slice can be nil
// slice has declaration and instantiation ..
// make is a builtin function that is used to instantiate a slice

/*
Slice Header
-----------
ptr: 8 bytes
len: 8 bytes
cap: 8 bytes
*/

/*
var slice1 []int
Slice Header
-----------
ptr:nil
len:0
cap:0
slice1=[]
*/

/*
slcie1 = make([]int,5)
Slice Header
-----------
ptr: 0x140000bc800
len: 5
cap: 5

// zero values are given to the slice, similar to array
slice1:=[0 0 0 0 0]
*/
