package main

import (
	"errors"
	"fmt"
)

func main() {
	var num = 100
	Sq(num)
	fmt.Println(num)
	err := SqP1(&num)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
	var ptr1 *int
	err = SqP1(ptr1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}

	ptr2 := new(int)
	*ptr2 = 25
	err = SqP1(ptr2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*ptr2)
	}

	ptr3, err := SqP2(&num)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*ptr3, num)
	}

}

func Sq(num int) {
	num = num * num
	//println(num)
}

func SqP1(num *int) error {
	if num != nil {
		*num = *num * *num
		return nil
	} else {
		return errors.New("nil pointer param")
	}
}

func SqP2(num *int) (*int, error) {
	if num != nil {
		*num = *num * *num
		return num, nil
	} else {
		return nil, errors.New("nil pointer param")
	}
}
