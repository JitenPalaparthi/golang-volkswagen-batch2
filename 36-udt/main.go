package main

import "fmt"

func main() {
	p1 := struct { // creating struct variable without struct name
		Id          int
		Name, Email string
		Status      string
		Greet       func(string)
	}{
		Id:     101,
		Name:   "Jiten",
		Email:  "JitenP@outlook.com",
		Status: "Active",
		Greet:  func(msg string) { fmt.Println(msg) },
	}
	fmt.Println(p1)
	p1.Greet("Hello VolksWagen")

	var p2 struct { // creating struct variable without struct name
		Id          int
		Name, Email string
		Status      string
	}
	fmt.Println(p2)
}
