package main

func main() {

	println("normal for loop")

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	println("1-20 even numbers")

	j := 1

	for j <= 20 { // for loop like a while loop
		if j%2 == 0 {
			print(j, " ")
		}
		j++
	}
	println()

	for i, j := 1, 10; i <= j && true && (true || false); i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	k := 1
	for {
		if k > 10 {
			break
		}
		println(k*k, " ")
		k++
	}

	l := 1

	for ; l <= 20; l++ {
		if l%2 == 0 {
			continue
		}
		print(l, " ")
	}
	println("\nnested loop")

out:
	for i := 1; i <= 10; i++ {
		for j := 2; j <= 10; j++ {
			if i == j {
				break out
			}
			println("i:", i, "j:", j)
		}
	}
}
