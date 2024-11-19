package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 5)                     // declared and instantiated with zero values
	slice2 := []int{10, 12, 32, 23, 4, 53, 5, 4} // shorthand with values
	fmt.Println(slice2)
	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	addressOf := fmt.Sprintf("%p", &slice1) // address of slice header
	fmt.Println("slice1:", slice1, "Address:", addressOf, "len:", len(slice1), "cap", cap(slice1), "Address of 0th", &slice1[0])
	slice1 = append(slice1, 999)
	fmt.Println("slice1:", slice1, "Address:", addressOf, "len:", len(slice1), "cap", cap(slice1), "Address of 0th", &slice1[0])
	slice1 = append(slice1, 9999)
	fmt.Println("slice1:", slice1, "Address:", addressOf, "len:", len(slice1), "cap", cap(slice1), "Address of 0th", &slice1[0])
	var slice3 []int

	slice3 = append(slice3, 10, 20, 30)
	fmt.Println("slice3:", slice3, "Address:", addressOf, "len:", len(slice3), "cap", cap(slice3), "Address of 0th", &slice3[0])
	slice3 = append(slice3, 40, 50, 60)
	fmt.Println("slice3:", slice3, "Address:", addressOf, "len:", len(slice3), "cap", cap(slice3), "Address of 0th", &slice3[0])
	slice3 = append(slice3, 70)
	fmt.Println("slice3:", slice3, "Address:", addressOf, "len:", len(slice3), "cap", cap(slice3), "Address of 0th", &slice3[0])

	slice4 := make([]int, 100)
	for i := range slice4 {
		slice4[i] = rand.Intn(1000)
	}
	for i := 1; i <= 100; i++ {
		slice4 = append(slice4, rand.Intn(1000))
	}
	//slice4 = append(slice4, rand.Intn(1000))

	fmt.Println("slice4:", slice4, "Address:", addressOf, "len:", len(slice4), "cap", cap(slice4), "Address of 0th", &slice4[0])
}

/*
if newCapacity < 1024:
    newCapacity = currentCapacity * 2
else:
    newCapacity = currentCapacity + currentCapacity / 4
*/

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
