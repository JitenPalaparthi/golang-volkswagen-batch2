package main

import "fmt"

type integer = int // integer is not a separate type , rather it is a same int with an alias

func main() {
	var num1 int       // 0
	var float1 float32 // 0
	var ok1 bool       // false
	var str1 string    // ""
	println(num1, float1, ok1, str1)
	num1 = 100
	float1 = 123.123
	ok1 = true
	str1 = "Hello World"
	fmt.Println(num1, float1, ok1, str1)
	var num2 integer = 12321
	var char1 rune = 'A'  // 1 byte
	var char2 int32 = '世' // 3 bytes
	var byte1 byte = 'A'  // 1 byte
	char2 = char2 + 10    // how can I perform arthemetic operation on a char
	fmt.Println(num2, byte1, char1, char2)
}

// numbers
// int, uint,int8,int16,int32,int64, uint8,uint16,uint32,uint64,float32,float64,

// rune, byte

// boolean
// bool

// strings
// string

// interface type
// interface{} or any

// there is a zero value for each variable in go
