package main

func main() {

	var num uint8 = 4

	switch num {
	case 1:
		println("sunday")
	case 2:
		println("monday")
	case 3:
		println("tuesday")
	case 4:
		println("wednesday")
	case 5:
		println("thursday")
	case 6:
		println("friday")
	case 7:
		println("satdurday")
	default:
		println("no day")
	}
	{
		var num int = 120
		//var k int
		switch { // empty switch
		case num >= 0 && num < 50:
			println(num, "is between 0-50")
		case num >= 50 && num < 100:
			println(num, "is between 50-100")
		case num >= 100 && num < 150:
			println(num, "is between 100-150")
			println("just to check")
		case num >= 150:
			println(num, "is 150 or greater")
		default:
			println(num, "is less then 0")
		}

	}

	char := '界'

	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is vovel")
	default:
		println(string(char), "is consonent or a special unicode character")
	}

}
