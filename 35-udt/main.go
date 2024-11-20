package main

import "fmt"

func main() {

	var v1 struct {
		Line1, City, State, Country, Pincode, Status string
	}
	fmt.Println(v1) // v1 prints zero values of the struct

	//var f1 func(int, int) int

	p1 := Person{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com", Status: "active", Address: struct {
		Line1, City, State, Country, Pincode, Status string
	}{Line1: "Some Line", City: "Bangalore", State: "karnataka", Country: "India", Pincode: "560086", Status: "active"}}

	fmt.Println(p1)

	p1.Print = func() {
		fmt.Println("Person Details")
		fmt.Println(p1)
	}
	p1.Print()

}

type Person struct {
	Id          int
	Name, Email string
	Status      string
	Address     struct {
		Line1, City, State, Country, Pincode, Status string
	} // embedded struct
	Print func()
}
