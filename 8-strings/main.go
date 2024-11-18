package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := "Hello World!" // 12 chars
	str2 := "你好,?-世界!$%&好" // 15 chars
	var str3 string

	fmt.Println("Size of str1:", unsafe.Sizeof(str1))
	fmt.Println("Size of str2:", unsafe.Sizeof(str2))
	fmt.Println("Size of str3:", unsafe.Sizeof(str3))
	// if str1 == nil {

	// }
	fmt.Println("first byte of str1:", &([]byte(str1)[0]))
	fmt.Println("address of str1:", &str1)

	str1 = "Hello Universe!" // am i mutating it or not?

	fmt.Println("first byte of str1:", &([]byte(str1)[0]))
	fmt.Println("address of str1:", &str1)
	fmt.Println(len(str2))

}

/* any1
ptr: 8 bytes 0x122122
type:8 bytes float64
*/

/*
str1
----
ptr:  0x1231213aab 8bytes
len: 8
*/
/*
str3
----
ptr: 8
len: 8
*/
/*
str3
----
ptr: 8
len: 8
*/
