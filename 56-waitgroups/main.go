package main

import (
	"fmt"
	"sync"
)

func main() {
	//var wg *sync.WaitGroup
	wg := new(sync.WaitGroup)

	wg.Add(14)
	//wg.Add(1) // counter++
	go func() {
		for i := 1; i <= 10; i++ {
			//wg.Add(1)
			go greet(wg, "Hello world-->"+fmt.Sprint(i))
		}
		wg.Done() // counter--
	}()

	//wg.Add(1)
	go func() {
		evenNumbers(10)
		wg.Done()
	}()

	//wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("End of Main, but will be crashed")
		wg.Done()
	}(wg)

	//wg.Add(1)
	go func(wg *sync.WaitGroup) {
		slice := []int{123, 123, 3, 2, 4, 34, 5, 4, 7, 6, 89, 56, 34, 65, 56, 657}
		s1 := sum(slice...)
		fmt.Println("Sum-------------->>>>", s1)
		wg.Done()
	}(wg)
	//wg.Add(1)
	//wg.Done()
	wg.Wait() // wait ? how long? untile the counter becomes 0
}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
func greet(wg *sync.WaitGroup, msg string) {
	//time.Sleep(time.Millisecond * 100)
	println(msg)
	wg.Done()
}

func evenNumbers(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			println("even->", i)
		}
	}
}
