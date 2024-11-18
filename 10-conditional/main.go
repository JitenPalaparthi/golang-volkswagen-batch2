package main

func main() {
	var age uint8 = 18
	if age >= 18 {
		println("eligible for vote, bcz age is ", age)
	} else {
		println("not eligible for vote, bcz age is ", age)
	}

	if num := 10; num%2 == 0 {
		println(num, "is even number")
	} else {
		println(num, "is odd number")
	}

	var (
		age2   uint8 = 29
		gender uint8 = 'M'
		//value        = 942342234234343434 //110100010011110111110100010011110000011000101101000000001010
	)

	//fmt.Printf("%b", value)

	if age2 >= 18 && (gender == 'F' || gender == 'f') {
		println("she is eligible for marriage in India bcz of age is", age2)
	} else if age2 >= 21 && (gender == 77 || gender == 'm') {
		println("he is eligible for marriage in India bcz of age is", age2)
	} else {
		println("not eligible for marriage")
	}
	//fmt.Println(gender)
	// if gender > 100 {

	// }

	//var n uint8 = 0b11010001
	//fmt.Println(string(value), string(n))
	// okay:= {
	// 	if true{
	// 		return true
	// 	}else{
	// 		return false
	// 	}
	// };

	char1 := "A" //17
	char2 := "A" //17

	var (
		char3 uint8 = 'A'
		char4 rune  = '界' // 0-255
	)

	if char1[0] == char2[0] {

	}

	if rune(char3) == char4 {

	}

}
