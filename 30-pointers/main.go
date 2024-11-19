package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 5)
	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	//var slice2 *[]int
	slice2 := &slice1 // both of the slices referred to the same memory
	fmt.Println("Slcie1:", slice1)
	for _, v := range *slice2 {
		fmt.Print(v, " ")
	}
	println()
	slice1 = append(slice1, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println("Slcie1:", slice1)
	// for _, v := range *slice2 {
	// 	fmt.Print(v, " ")
	// }
	fmt.Println("slice2:", *slice2)
	fmt.Printf("address of slice1:%p address of slice2:%p", &slice1, slice2)
	(*slice2)[0] = 9999
	fmt.Println("slice2:", *slice2)

}
