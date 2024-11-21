package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 1; ; i++ {
			if i > 100 {
				runtime.Goexit() // in which goroutine?
				//return
			}
			go greet("Hello world-->" + fmt.Sprint(i))
		}
	}()
	go func() {
		evenNumbers(1000)
		runtime.Goexit()
	}()
	go fmt.Println("End of Main, but will be crashed")

	go func() {
		slice := []int{123, 123, 3, 2, 4, 34, 5, 4, 7, 6, 89, 56, 34, 65, 56, 657}
		s1 := sum(slice...)
		fmt.Println("Sum-------------->>>>", s1)
	}()

	time.Sleep(time.Second * 1)
}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
func greet(msg string) {
	time.Sleep(time.Millisecond * 100)
	println(msg)
}

func evenNumbers(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			go println("even->", i)
		}
	}
}
