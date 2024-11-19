package main

import (
	"math/rand"
)

func main() {

	slice1 := make([]int, 5)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 1, 2, 3, 4, 5

	for _, v := range slice1 {
		print(v, " ")
	}
	println()

	var any1 any = slice1

	for _, v := range any1.([]int) {
		print(v, " ")
	}

	slice2 := make([]int, 10000)
	for i := range slice2 {
		slice2[i] = rand.Intn(100000)
	}
	slice1 = append(slice1, slice2...)
	for i := 0; i < 10; i++ {
		print(slice1[i], " ")
	}
	any1 = slice2

	// var arr1 [1000000000]int
	// println(arr1[0])
	println("-------------------")
	// ---------------------------

	var num1 = 100
	ptr1 := Sq(&num1)
	println(*ptr1)
	println("address of num1:", &num1)
	println("address that  ptr1:", ptr1)

	var num2 = 4
	ptr2 := SqP(&num2)
	println(*ptr2)
	println("address of num2:", &num2)
	println("address that  ptr2:", ptr2)

	var num3 = 6
	any2 := SqAny(&num3)
	println(any2)

}

func Sq(num *int) *int {
	if num != nil {
		*num = *num * *num
	}
	return num
}

func SqP(num *int) *int {
	ret := new(int) // pointer is created inside a function
	if num != nil {
		*ret = *num * *num
	}
	println("address of inside:", ret)
	return ret // leacked outside of the function , the problem is dangling pointer
}

func SqAny(num *int) any {
	//ret := new(int) // pointer is created inside a function
	var ret any // any internally contains a pointers
	if num != nil {
		ret = *num * *num
	}
	//println("address of inside:", ret)
	return ret // leacked outside of the function , the problem is dangling pointer
}

// does not escape means that variable has not created in heap memory
// escapes to heap means it has been stored in heap memory
// leaking param: num to result ~r0 level=0

// go build -gcflags="-m" main.go
