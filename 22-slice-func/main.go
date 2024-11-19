package main

import (
	"math/rand"
)

func main() {
	slice1 := make([]int, 5) // declared and instantiated with zero values

	for i := range slice1 {
		slice1[i] = rand.Intn(10)
	}
	GetSliceHeader(slice1)
	println("Before double:", slice1)
	slice1 = DoubleSlice(slice1, 11, 12)
	//slice2 := DoubleSlice(slice1, 11, 12)
	println("After double:", slice1)
	//fmt.Printf("Address of slice1:%p\n", &slice1)
	GetSliceHeader(slice1)
	// slice3 := make([]string, 1000000)
	// println("Size of slice3:", unsafe.Sizeof(slice3))
}

func GetSliceHeader(slice []int) {
	if slice != nil {
		println("Address:", &slice)
		println("Slice Header")
		println("-------------")
		println("Pointer:", &slice[0])
		println("Len:", len(slice))
		println("Cap:", cap(slice))
		println("-------------")

	}
}

func DoubleSlice(slice []int, a, b int) []int {
	//fmt.Printf("Address Of:%p\n", &slice)
	//slice = append(slice, a, b) // as soon as the append is used here it became divegetn
	for i, v := range slice {
		slice[i] = v * v
	}
	GetSliceHeader(slice)

	return slice
}

/*
before append
slice header
-----------
ptr: 0x14000016090
len: 5
cap: 5
*/

/*
after append
slice header
-----------
ptr: 0x1400001a0f0
len: 6
cap: 10
*/

/*
Slice Header
-----------
ptr: 8 bytes
len: 8 bytes
cap: 8 bytes
*/
