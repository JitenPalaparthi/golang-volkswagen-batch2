package main

import "time"

func main() {
	defer println("End of main")
	ch := make(chan int)
	//signal := make(chan bool) // bool -> 1 byte
	signal := make(chan struct{}) // bool -> 1 byte
	//fmt.Println("size of", unsafe.Sizeof(ch))
	num := 10
	go sendSq(ch, num)
	go receiveSq(ch, signal, num)

	// v := <-signal // until the value is received
	// fmt.Println(v)
	<-signal
}

func sendSq(ch chan int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i * i
	}
}

func receiveSq(ch chan int, sig chan struct{}, num int) {
	for i := 1; i <= num; i++ {
		time.Sleep(time.Millisecond * 100)
		println(<-ch)
	}

	sig <- struct{}{}
}
