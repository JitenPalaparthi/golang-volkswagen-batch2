package main

import (
	"fmt"
)

func main() {

	var any1 any = true

	switch any1.(type) {
	case int:
		v := any1.(int)
		println(v * v)
	case uint8:
		v := any1.(uint8)
		println(v * v)
	case float32:
		v := any1.(float32)
		fmt.Printf("%.2f", v*v)
	case float64:
		println("jumps")
		v := any1.(float64)
		fmt.Printf("%.2f", v*v)
	case string:
		println("cannot perform arthemetic on string")
	case bool:
		println("cannot perform arthemetic on bool")

	default:
		println("undefined type, not in the given types")
	}

	any1 = "Hello World"

	switch v1 := any1.(type) {
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		println(v1)
	default:
		println("not a number type")
	}

	// any1 = 12.123
	// val := reflect.ValueOf(any1)
	// println(val.Float() * val.Float())

}

// whereever you remove break in other programmin languages, add fallthrough in golang
