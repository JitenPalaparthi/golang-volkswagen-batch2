package main

import (
	"fmt"
)

func main() {
	p1 := Person{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com", Status: "active", Address: &Address{Line1: "Some Line", City: "Bangalore", State: "karnataka", Country: "India", Pincode: "560086", Status: "active"}}
	fmt.Println(p1)
	fmt.Println("pincode", p1.Address.Pincode)

	p1.Address.PrintAddress()

	s1 := Student{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com", Status: "active", Address: &Address{Line1: "Some Line", City: "Bangalore", State: "karnataka", Country: "India", Pincode: "560086", Status: "active"}}

	fmt.Println("pincode:", s1.Pincode)
	s1.PrintAddress()
	fmt.Println("status of student", s1.Status)
	fmt.Println("status of student Address", s1.Address.Status)

}

type Address struct {
	Line1, City, State, Country, Pincode, Status string
}

func (a *Address) PrintAddress() {
	fmt.Println("Line1:", a.Line1)
	fmt.Println("City:", a.City)
	fmt.Println("State:", a.State)
	fmt.Println("Country:", a.Country)
	fmt.Println("Pincode:", a.Pincode)
	fmt.Println("Status:", a.Status)
}

type Person struct {
	Id          int
	Name, Email string
	Status      string
	Address     *Address // composition
}

type Student struct {
	Id          int
	Name, Email string
	Status      string
	*Address    // promoted field
}

// // example GORM pacakge
// type user struct {
// 	gorm.Model
// 	Name, Email string
// }

// // Id, create_at, delted_at,updated_at
// u1.ID
// u1.create_at
// u1.Name
