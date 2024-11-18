package main

import "fmt"

func main() {

	var num1 uint8 = 127 // 1 byte

	var num2 uint = uint(num1) // 8 bytes
	// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 01111111

	fmt.Printf("num2 value: %d and num2 type:%T\n", num2, num2)

	var num3 uint32 = 1231231 // 4 bytes
	//100101100100101111111

	fmt.Printf("num3 binary:%b\n", num3)

	var num4 uint8 = uint8(num3) // 01111111

	//var num5 = 0b01111111
	//fmt.Println(num5)

	fmt.Printf("num4 value: %d and num4 type:%T", num4, num4)

	var float1 = 1231.12312

	var num5 uint = uint(float1)

	var rune1 rune = 'K'

	var num6 uint16 = uint16(rune1)
	//num6 = (uint16)(rune1)

	var ok1 bool = true // 1 or 0

	// var byte1 byte = byte(ok1) // this cannot be casted
	fmt.Println(num5, string(rune1), num6, ok1)

	byte2 := byte(65)
	num7 := uint16(19000)
	num8 := uint64(19000)

	var num9 float32 = 65.123

	fmt.Println("byte2 to string:", string(byte2))
	fmt.Println("num7 to string:", string(num7))
	fmt.Println("num7 to string:", string(num8))
	fmt.Println("num9 to string:", string(uint(num9)))

	// complex numbers
	// complex64 and complex128

	c1 := complex(123.123, 34.4)
	r1, im1 := float32(123.123), float32(34.4)
	c2 := complex(r1, im1)
	c3 := complex(float32(123.123), float32(34.4))
	c4 := 123.23 + 5.2i // r+mi it is a complex number

	c5 := c1 + complex128(c2)

	c6 := complex128(c1) - c4

	c7 := complex128(c3) * c4 // mathematically done
	c8 := complex128(c3) / c4 // mathematically done

	fmt.Println(c5, c6, c7, c8)

	r2, im2 := real(c8), imag(c8)
	fmt.Println("real r2:", r2, "imag im2:", im2)
}

// there is no implicit type casting in go
// every type is distinct
// a variable from one type can only be assigned to anoter type using cast or conversion
// all numbers can be casted to one type to another type
