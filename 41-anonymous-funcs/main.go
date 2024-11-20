package main

import "fmt"

func main() {

	a, b := 10, 20

	sum1 := calc(a, b, func(i1, i2 int) int {
		return i1 + i2
	})

	// sub1 := calc(a, b, func(i1, i2 int) int {
	// 	return i1 - i2
	// })
	sub1 := calc(a, b, sub)

	// mul1 := calc(a, b, func(i1, i2 int) int {
	// 	return i1 * i2
	// })

	m1 := func(i1, i2 int) int {
		return i1 * i2
	}

	mul1 := calc(a, b, m1)

	fmt.Println("sum1:", sum1)
	fmt.Println("sub1:", sub1)
	fmt.Println("mul1:", mul1)

	num := 1

	v := func() int {
		for i := 1; i <= 10; i++ {
			num++
		}
		return num
	}()

	fmt.Println("num1:", v)
	fmt.Println("num1:", num)

	//map1 := make(map[string]func(int, int) int)
	map1 := make(map[string]any)

	map1["add"] = func(a, b int) int {
		return a + b
	}
	map1["sub"] = func(a, b int) int {
		return a - b
	}
	map1["mul"] = func(a, b int) int {
		return a * b
	}
	map1["dib"] = func(a, b int) int {
		return a / b
	}
	map1["rem"] = func(a, b int) int {
		return a % b
	}
	map1["square"] = func(s int) int {
		return s * s
	}
	map1["value"] = 100

	map1["greet"] = "hello world"

	for k, v := range map1 {
		// r := v.(func(int, int) int)(10, 20)
		// fmt.Println(k, "->", r)
		switch v := v.(type) {
		case func(int, int) int:
			r := v(10, 20)
			fmt.Println(k, "->", r)
		case func(int) int:
			r := v(100)
			fmt.Println(k, "->", r)
		case string:
			fmt.Println(v)
		default:
			fmt.Println("un implemented")
		}
	}

}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func sub(a, b int) int {
	return a - b
}
