package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err.Error())
	}

	bytes := make([]byte, 10)
	for {
		_, err := file.Read(bytes)
		if err != nil {
			break
		}
		fmt.Print(string(bytes))
	}
	file.Close()
}
