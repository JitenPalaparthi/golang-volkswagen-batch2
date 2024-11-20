package main

import (
	"math/rand"
)

func main() {

	slice1 := make([]int, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100)
	}

	// for i := 0; i <= len(slice1); i++ {
	// 	fmt.Print(slice1[i], " ")
	// }

	// num := 0

	// println(100 / num)

	var ptr1 *int

	*ptr1 = 100

	// comment and uncomment the each section of the code to see different kinds of panics

}
