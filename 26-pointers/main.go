package main

import "fmt"

func main() {

	var num1 int = 100

	var ptr1 *int = &num1

	//ptr1 = nil

	var num2 = *ptr1

	// if *ptr1 >= 100 {

	// }

	//*ptr1++ // cannot dereference a nil pointer

	fmt.Println(num1, num2)

	var ptr2 *int // zero value of a pointer is nil

	if ptr2 == nil {
		fmt.Println("ptr2 is nil")
		ptr2 = new(int) // in allocates some memory to store int values and zero values are given  to the poiners
	}
	fmt.Println("ptr2", ptr2, "*ptr2", *ptr2)

	ptr3 := new(bool)
	fmt.Println("ptr3", ptr3, "*ptr3", *ptr3)

	*ptr3 = true
	fmt.Println("ptr3", ptr3, "*ptr3", *ptr3)
	var ptr4 **int = &ptr1
	**ptr4 = 500
	fmt.Println(num1)
	var ptr5 ***int = &ptr4
	***ptr5 = 999
	fmt.Println(num1)
}
