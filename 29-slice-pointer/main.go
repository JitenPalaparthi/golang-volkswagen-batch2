package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"unsafe"
)

func main() {
	slice1 := make([]int, 5)
	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	fmt.Println(slice1)
	ptr := uintptr(unsafe.Pointer(&slice1[0])) // first element addrtess
	val := (*int)(unsafe.Pointer(ptr))
	fmt.Println(*val)
	size := unsafe.Sizeof(slice1[0]) // 8 bytes
	for i := 0; i < len(slice1)-1; i++ {
		ptr = ptr + size
		val := (*int)(unsafe.Pointer(ptr))
		fmt.Println(*val)
	}

	//reflect.SliceHeader
	//reflect.StringHeader

	slice := []int{10, 20, 30, 40, 50}

	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Println("pointer:", sliceheader.Data)
	fmt.Println("Len:", sliceheader.Len)
	fmt.Println("Cap:", sliceheader.Cap)

	ptr2 := sliceheader.Data
	ptr2 += 8
	val1 := (*int)(unsafe.Pointer(ptr2))
	fmt.Println(*val1)

	// var any1 any = 100

	// ptr3 := (*[2]any)(unsafe.Pointer(&any1))

	// fmt.Println((*ptr3)[0])
	// fmt.Println((*ptr3)[1])

}
