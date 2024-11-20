package main

import (
	"fmt"
	"reflect"
)

type integer int // creating a new type from the existing type

func (i *integer) ToString() string {
	return fmt.Sprint(*i)
}

func (i *integer) ToBytes() []byte {
	return []byte(fmt.Sprint(*i)) // though not a correct approach
}

type mybool bool

func (mb mybool) ToInt() uint8 {
	if mb == true {
		return 1
	}
	return 0
}

func (mb mybool) ToString() string {
	if mb == true {
		return "true"
	}
	return "false"
}

func main() {

	var num1 integer = 12334
	str1 := num1.ToString()
	fmt.Println("integer to string:", str1, "type of str1", reflect.TypeOf(str1))
	bytes1 := num1.ToBytes()
	fmt.Println("integer to bytes:", bytes1, "type of bytes1", reflect.TypeOf(bytes1))

	var num2 int = 12312

	str2 := (*integer)(&num2).ToString()
	fmt.Println("integer to string:", str2, "type of str2", reflect.TypeOf(str2))

	var float1 float32 = 12312.1232
	num3 := int(float1)
	str3 := (*integer)(&num3).ToString()
	fmt.Println("float32 to string:", str3, "type of str3", reflect.TypeOf(str3))

}
