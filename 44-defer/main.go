package main

import "os"

func main() {
	defer func() {
		println("End of main-1")
	}()
	defer func() {
		println("End of main-2")
	}()
	defer func() {
		println("End of main-3")
	}()

	// file, err := os.OpenFile("data1.txt", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer file.Close()
	writeFile("data1.txt", "Hello folks") // change the filename to see the panic and examine differes

	func() {
		for i := 1; i <= 20; i++ {
			if i%2 != 0 {
				continue
			}
			print(i, " ")
		}
	}()

}

func Fatal(msg string) {
	println(msg)
	os.Exit(1) // exits the application abruptly
}

func writeFile(filename, msg string) error {
	defer func() {
		println("odd numbers")
		for i := 1; i <= 20; i++ {
			if i%2 != 0 {
				print(i, " ")
			}
		}
	}()

	defer println("Will I be executed")

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	_, err = file.WriteString(msg)

	return err
}
