package main

import "fmt"

func main() {
	//fmt.Println("hello", 100, true, 123.123) // fmt.Println contains variadic argument
	sum1 := SumOf()
	println("sum1:", sum1)
	sum2 := SumOf(10, 20)
	println("sum2:", sum2)
	sum3 := SumOf(10, 123, 123, 4, 34, 5, 35, 6, 5, 8, 68, 64, 45, 5, 46, 4, 34, 34)
	println("sum3:", sum3)
	msum1 := MultiplySumOf(2, 1, 2, 3, 4, 5, 6)
	println("msum1:", msum1)

	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	sum4 := SumOf(slice1...) // a slice can be passed  as a variadic argument using slice...
	println("sum4:", sum4)

	arr1 := [3]int{1, 2, 3}

	slice2 := arr1[:] // an array can be converted as alice
	fmt.Println("slice2:", slice2)
	slice2 = append(slice2, 4, 5, 6)
	slice2[1] = 9999
	fmt.Println("slice2:", slice2)
	fmt.Println("arr1:", arr1)

	sum5 := SumOf(arr1[:]...)
	println("sum5:", sum5)

	slice4 := make([]int, 3)

	copy(slice4, arr1[:])

}

//	func SumOf(a, b int) int {
//		return a + b
//	}
//func MultiplySumOf(nums ...int, n int) int { //not okay .. variadic param must be the last argument in a function

func MultiplySumOf(n int, nums ...int) int { // okay since variadic is last param
	sum := 0
	for _, v := range nums {
		sum += v * n
	}
	return sum
	//return a + b
}

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
	//return a + b
}

// variadic parameter ...T
// variadic parameter should be the last parameter in a function/method
// cannot create variadic as a filed , or a normal variable
// can use range loop for variadic like on a array or slice

func SumOfAny(nums ...any) float64 {
	// only numbers to be added, is bool , true add 1 or false add 0,
	//  if string it contain a number then add that number if not just leave it
	// sum := 0
	// for _, v := range nums {
	// 	sum += v
	// }
	// return sum
	return 0
	//return a + b
}
