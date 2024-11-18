package main

import "fmt"

func main() {

	var (
		num1   int
		num2   uint8
		num3   int32
		num4   int64
		num5   uint
		float1 float32
		float2 float64
		rune1  rune // alias for int32
		byte1  byte // alias for uint8
	)

	fmt.Println(num1, num2, num3, num4, num5, float1, float2, rune1, byte1)

	var num6 = 12312     // type inference, inferred as int
	var float3 = 123.123 // type inferrence , inferred as float64
	var age = 39         // 8 bytes in 64bit arch

	var ok1 = true           // type is inferred as bool , it is 1 byte
	var str1 = "Hello World" // type is inferred as string , which is 16 bytes with a allocation
	var char1 = 'A'          // rune is used which is int32

	fmt.Println(num6, float3, age, ok1, str1, char1)

	// short hand declaraion of a variable
	a := 100   // a is inffered as int , owning the value 100
	b := false // inffered as boolen
	c := 123.123
	str2 := "Hello VW Minds"
	fmt.Println(a, b, c, str2)

	d, e, f, g := 100, false, 123.123, "VolksWagen Corp" // shorthand declaration with type inferrence
	fmt.Println(d, e, f, g)

	// no direct shadowing
	//d := 300.1231 // can't do in the same scope
	{
		// can shadow in a separate scope
		d := 300.1231 // shadowing
		fmt.Println(d)
	}
}
