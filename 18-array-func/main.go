package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr1 [10]uint16
	for i := 0; i < len(arr1); i++ {
		arr1[i] = uint16(rand.Intn(100))
	}
	fmt.Println(arr1)
	sum1 := SumOf(arr1)
	fmt.Println("Sum of arr1:", sum1)

	// arr2 := [5]uint16{10, 30, 4, 5, 60}
	// sum2 := sumOf(arr2)
}

func SumOf(arr [10]uint16) int {
	sum := 0

	for _, v := range arr {
		sum += int(v)
	}
	return sum
}

//const arr = [3]int{1, 2, 3} // not allowed , cant create const as an array
