package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 5, 10) // declared and instantiated with zero values

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}

	slice2 := slice1 // shallow copy , headers are copied

	fmt.Println("Slice1")
	GetSliceHeader(slice1)
	fmt.Println("Slice1")
	GetSliceHeader(slice2)
	slice2[0] = 12312
	fmt.Println("slice1:", slice1)
	slice1 = append(slice1, 999) // both slice1 and slice2 are diverged
	//slice2 = slice1
	slice2[1] = 101010
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

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
