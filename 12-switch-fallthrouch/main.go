package main

func main() {

	var num int = 24

	// fallthrough in a positive way
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("------------------------------")
	println("false negative")

	num = 6
	// fallthrough in a positive way
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

}

// whereever you remove break in other programmin languages, add fallthrough in golang
