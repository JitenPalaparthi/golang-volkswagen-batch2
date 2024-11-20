package main

func main() {

	a := 100

	defer println("\ndefer print of a:", a)

	a *= a

	println("normal print of a:", a)

	str := "Hello World"
	//println()

	for _, v := range str {
		defer print(string(v), "")
	}

	//println()

	b := new(int)
	*b = 100
	func(ptr *int) {
		defer println("defer print of b:", *b)
		*ptr = *ptr + 200
	}(b)
	println("normal print of b:", *b)

}
