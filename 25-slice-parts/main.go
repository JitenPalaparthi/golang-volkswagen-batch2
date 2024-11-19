package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 20)
	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	fmt.Println("slice1:", slice1)
	slice2 := slice1 // the whole slice is headers copied
	fmt.Println("slice2:", slice2)
	slice3 := slice1[:5] // 0 ,1, 2, 3,4 indices are taken into slice3
	fmt.Println("slice3:", slice3)
	slice4 := slice1[5:] // 5th element to the end
	fmt.Println("slice4:", slice4)
	slice5 := slice1[10:15] // 10th to 14th but not 15th
	fmt.Println("slice5:", slice5)

	slice5[0] = 9999999
	fmt.Println("slice1:", slice1)
	slice6, err := RemoveElem(slice1, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("slice6:", slice6)

	}

	var slice7 []int

	slice8, err := RemoveElem(slice7, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("slice8:", slice8)
	}

	slice9, err := RemoveElem(slice1, 17)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("slice9:", slice9)
	}

}

func RemoveElem(slice []int, i int) ([]int, error) {
	if slice == nil {
		return nil, errors.New("nil slice")
	}
	if i >= len(slice) {
		return nil, fmt.Errorf("%d is out of bounds of the slice", i)
	}
	// [93 10 62 61 17 29 8 45 79 88 16 37 52 16 98]
	// 93 10 62 61 17 [8 45 79 88 16 37 52 16 98]...
	slice = append(slice[:i], slice[i+1:]...)

	return slice[:len(slice)-1], nil
}
