package main

import (
	_ "time"
)

func main() {

	str1 := "Hello, World!"

	for i := 0; i < len(str1); i++ {
		print(str1[i], "->", string(str1[i]), " ")
	}
	println()
	str2 := "Hello, 世界!"

	for i := 0; i < len(str2); i++ {
		print(str2[i], "->", string(str2[i]), " ")
	}
	println()
	// _ blank identifier
	for _, v := range str2 {
		print(string(v), " ")
	}

}
