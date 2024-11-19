package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 5) // declared and instantiated with zero values

	for i := range slice1 {
		slice1[i] = rand.Intn(10)
	}
	slice2 := make([]int, 5)
	copy(slice2, slice1) //deep copy
	slice3 := make([]int, 3)
	copy(slice3, slice2)

	fmt.Println(slice1, slice2, slice3)
	clear(slice1)
	fmt.Println(slice1)
}

func GetSliceHeader(slice []int) {
	if slice != nil {
		fmt.Printf("Address:%p\n", &slice)
		fmt.Println("Slice Header")
		fmt.Println("-------------")
		fmt.Println("Pointer:", &slice[0])
		fmt.Println("Len:", len(slice))
		fmt.Println("Cap:", cap(slice))
		fmt.Println("-------------")

	}
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
