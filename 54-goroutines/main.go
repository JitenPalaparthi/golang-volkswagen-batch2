package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello World")
	}()
	go Greet()
	fmt.Println("End of main")
	time.Sleep(time.Second * 3)
	//runtime.Goexit()
}

func Greet() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello VW folks")
		time.Sleep(time.Millisecond * 100)
	}
	runtime.Goexit()
}

// 1. main is a goroutune
// 2. no go routine waits for other goroutine to completes its execxution
// 3. order of goroutine execution is premediated
// 4. to create a goroutine just use go keyword

// what is a thread?
// What happens when a thread is created?
// what are system threads and green threads?
// what is context switching?
// what happen to values during context switch?
// what happen when threads are blocked

// 12 core    12   hardware threads N
// 674 total  3663 threads          M
// N : M
// goscheduler
// go does context switching
// go creates processes (threads with some additional data)
// number of processes are by default equal to number of hardware cores in the system.
