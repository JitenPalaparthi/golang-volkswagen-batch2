package main

import "fmt"

func main() {
	defer recoverThis()
	println("Start of main")
	defer println("End of main")

	FullName("Jiten", "P")
	FullName("Rahim", "")
	EvenNumbers(20)

	num := 0
	println(100 / num)

}

func recoverThis() {
	if v := recover(); v != nil {
		fmt.Println("recovering it:", v)
	}
}

func FullName(firstname, lastname string) {
	defer func() {
		println("Odd numbers")
		for i := 1; i <= 10; i += 2 {
			print(i, " ")
		}
		println()
	}()
	defer recoverThis()
	if firstname == "" {
		panic("firstname is empty")
	}
	if lastname == "" {
		panic("lastname is empty")
	}
	println(firstname + " " + lastname)
}

func EvenNumbers(num int) {
	i := 1
loop:
	if i%2 == 0 {
		print(i, " ")
	}
	i++
	if i <= num {
		goto loop
	}
}
