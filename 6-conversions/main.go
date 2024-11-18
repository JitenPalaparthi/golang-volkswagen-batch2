package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var str1 = "1234"
	num, err := strconv.Atoi(str1) // 1. a function can return multiple values 2. err is just a value

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Value of num:", num, "Type of num:", reflect.TypeOf(num))

	var str2 = "1234a"
	num, err = strconv.Atoi(str2) // 1. a function can return multiple values 2. err is just a value

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value of num:", num, "Type of num:", reflect.TypeOf(num))
	}

	str3 := "1234.2342"

	float1, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value of float1:", float1, "Type of float1:", reflect.TypeOf(float1))
	}
	//1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False

	str4 := "1"
	ok, err := strconv.ParseBool(str4)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value of ok:", ok, "Type of ok:", reflect.TypeOf(ok))
	}

	var num1 = 65

	str5 := strconv.Itoa(num1) // as it is converted as string
	fmt.Println("Value of str5:", str5, "Type of str5:", reflect.TypeOf(str5))

	str6 := string(num1) // type casted as A
	fmt.Println("Value of str6:", str6, "Type of str6:", reflect.TypeOf(str6))

	// var num3 uint8 = 123
	// var num4 uint64 = uint64(num3)
	//a, b := 10, 20
}
