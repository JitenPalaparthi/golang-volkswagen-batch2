package main

import (
	"fmt"
)

type integer int // creating a new type from the existing type

func (i integer) ToString() string {
	return fmt.Sprint(i)
}

type myint integer

func (i myint) ToBytes() []byte {
	return []byte(fmt.Sprint(i)) // though not a correct approach
}

type mybool bool

func (mb mybool) ToInt() uint8 {
	if mb == true {
		return 1
	}
	return 0
}

func (mb mybool) ToString() string {
	if mb {
		return "true"
	}
	return "false"
}

func main() {

	var num1 int = 12321

	var num2 integer = 342323

	var num3 myint = 123123

	str1 := integer(num1).ToString()
	//bytes1 := myint(integer(num1)).ToBytes() //okay
	bytes1 := myint(num1).ToBytes()

	fmt.Println("num1 to str:", str1)
	fmt.Println("num1 to bytes:", bytes1)

	str2 := num2.ToString()
	bytes2 := myint(num2).ToBytes()

	fmt.Println("num2 to str:", str2)
	fmt.Println("num2 to bytes:", bytes2)

	str3 := integer(num3).ToString()
	bytes3 := num3.ToBytes()

	fmt.Println("num3 to str:", str3)
	fmt.Println("num3 to bytes:", bytes3)

	var float1 float64 = 123213.123123

	str4 := integer(float1).ToString()
	bytes4 := myint(float1).ToBytes()
	fmt.Println("float1 to str:", str4)
	fmt.Println("float1 to bytes:", bytes4)

	var ok1 mybool = true

	str5 := integer(ok1.ToInt()).ToString()
	bytes5 := myint(ok1.ToInt()).ToBytes()

	fmt.Println("ok1 to str:", str5)
	fmt.Println("ok1 to bytes:", bytes5)

}
