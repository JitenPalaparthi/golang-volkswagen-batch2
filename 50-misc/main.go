package main

import "fmt"

func main() {
	a, b := 10, 20
	// t := a
	// a = b
	// b = t
	fmt.Println("before swap a,b", a, b)
	a, b = b, a // swap
	fmt.Println("after swap a,b", a, b)
	a, b, c := 10, 20, 30
	fmt.Println("before swap a,b,c", a, b, c)
	a, b, c = c, a, b
	fmt.Println("after swap a,b,c", a, b, c)

	// var m map[string]any
	// if m == nil {
	// 	m = make(map[string]any)
	// }
	// m["name"] = "Jiten"
}
