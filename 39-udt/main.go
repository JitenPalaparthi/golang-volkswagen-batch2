package main

import (
	"errors"
	"fmt"
)

type mymap map[string]any

// this only init the this copy of the map which is not a pointer
func (m mymap) Init() {
	if m == nil {
		m = make(mymap)
	}
}

func New() mymap {
	var m mymap
	if m == nil {
		m = make(mymap)
	}
	return m
}

func (m mymap) Getkeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	l := len(m)
	keys := make([]string, l)

	c := 0
	for k := range m {
		keys[c] = k
		c++
	}
	return keys, nil
}

func (m mymap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	l := len(m)
	values := make([]any, l)

	c := 0
	for _, v := range m {
		values[c] = v
		c++
	}
	return values, nil
}

func (m mymap) Delete(key string) error {
	if m == nil {
		return errors.New("nil map")
	}

	_, ok := m[key]
	if !ok {
		return errors.New("no key found")
	}
	delete(m, key)
	return nil
}

func main() {

	var m1 mymap

	m2 := make(map[string]any)

	m2["560086"] = "Bangalore-1"
	m2["560096"] = "Bangalore-2"
	m2["560096"] = "Bangalore-2"
	m2["522001"] = "Guntur-1"
	m2["522002"] = "Guntur-2"

	m1 = New()

	m1["560086"] = "Bangalore-1"
	m1["560096"] = "Bangalore-2"
	m1["560096"] = "Bangalore-2"
	m1["522001"] = "Guntur-1"
	m1["522002"] = "Guntur-2"

	keys, err := m1.Getkeys()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Keys of a map:", keys)
	}

	values, err := m1.GetValues()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values of a map:", values)
	}

	err = m1.Delete("522002")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("key successfully deleted")
	}

	err = m1.Delete("522002")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("key successfully deleted")
	}

	keys, err = mymap(m2).Getkeys()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Keys of a map m2:", keys)
	}

}
