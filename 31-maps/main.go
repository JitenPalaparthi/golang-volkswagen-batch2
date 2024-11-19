package main

import "fmt"

func main() {

	var map1 map[string]any // this is only a declaration
	if map1 == nil {
		map1 = make(map[string]any)
	}

	map1["bangalore-1"] = "yeshvantpur and rajani nagar"
	map1["bangalore-2"] = "BTM Layout"
	map1["bangalore-3"] = "Whitefield"
	map1["bangalore-4"] = "Yelahanka"
	map1["bangalore-5"] = 560046

	val, ok := map1["bangalore-3"]
	if !ok {
		fmt.Println("key not available")
	} else {
		fmt.Println(val)
	}

	for k, v := range map1 { // for map range loop gives key and value
		fmt.Println("Key:", k, "Value:", v)
	}

	delete(map1, "bangalore-3")
	fmt.Println("after delete")
	for k, v := range map1 { // for map range loop gives key and value
		fmt.Println("Key:", k, "Value:", v)
	}

	val, ok = map1["bangalore-3"]
	if !ok {
		fmt.Println("key not available")
	} else {
		fmt.Println(val)
	}

}

// hashfunc
// 8 buckets , a bucket could be an array or a linkedlist
// echo "bangalore-1" | shasum
// cd8570dc63d6e8b7cc1a6bfdec78448a4b010c5b
// echo "bangalore-1 " | shasum
// c10f5efc4e802a1cbfaf87985d4b9f756951bf51
// cd8570dc63d6e8b7cc1 a6bfdec78448a4b010c5b
//  4                   2

// make is used to intantiate a map
// map is not ordered
// hashfunc is applied so that map is very efficient O(1) constant times fetch
// map is not thread safe

// what can be the key

// key is any datatype where == can be applied
// value can be anything
