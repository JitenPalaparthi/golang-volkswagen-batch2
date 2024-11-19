package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var arr1 [5]int // type of arr1 : [5]int
	// var arr2 [5]int  // type of arr1 : [5]int
	// var arr3 [10]int // type of arr2 : [10]int

	// if arr1 == arr3 { // cant do this since comparision operator cant be used on two different types

	// }

	//fmt.Println(arr1, arr2)

	fmt.Println("Type of arr1", reflect.TypeOf(arr1))
	fmt.Println("Size of arr1:", unsafe.Sizeof(arr1))
	fmt.Printf("Address of arr1:%p\n", &arr1)
	fmt.Printf("Address of arr1:%p", &arr1[0])

	arr1[0], arr1[1], arr1[2], arr1[3], arr1[4] = 100, 200, 300, 400, 500
	fmt.Println(arr1)
	sum := 0
	// for i := 0; i < len(arr1); i++ {
	// 	sum = sum + arr1[i]
	// sum+=arr[i]
	// }

	for _, v := range arr1 {
		sum += v
	}

	fmt.Println("Sum of arr1:", sum)

	// shorthand declartion of array
	arr2 := [5]int{10, 11, 12, 13, 14}
	//arr3 := [...]int{23, 2133, 4, 23, 45, 6, 56, 95, 45, 64, 23} // length is evaluated at compile time

	for i := len(arr2) - 1; i >= 0; i-- {
		print(arr2[i], " ")
	}

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	arr4d := [...][2][2][2]int{{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}, {{{9, 10}, {11, 12}}, {{13, 14}, {15, 16}}}}
	fmt.Println(arr3d, arr4d)

	println("iterating 2d array")
	for i1, a1 := range arr2d {
		fmt.Println("Type of a1:", reflect.TypeOf(a1), "\n")
		for i2, a2 := range a1 {
			print(a2, " ")
			a2 = a2 * a2
			arr2d[i1][i2] = a2
		}
	}

	println()
	fmt.Println(arr2d)

}

// arrays
// arrays are fixed length, the length should be knwon to the compiler
// zero value for the elements of an array
// what is the type of this  arr1 array? the type or an array includes its length
// two different arrays cannto be type casted or converted directly
