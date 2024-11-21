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
	go receiveSq(ch, signal)

	// v := <-signal // until the value is received
	// fmt.Println(v)
	<-signal //simialr to future
}

func sendSq(ch chan<- int, n int) {
	for i := 1; i <= n; i++ {
		// i := 1
		// for {
		ch <- i * i
		//}
	}
	close(ch) // close channel also sends a signal , will show later in select
}

func receiveSq(ch <-chan int, sig chan struct{}) {
	// for {
	// 	time.Sleep(time.Millisecond * 1000)
	// 	v, ok := <-ch // when the channel is closed , ok is false
	// 	if ok {
	// 		println(v)
	// 	} else {
	// 		sig <- struct{}{}
	// 		runtime.Goexit()
	// 	}
	// }

	for v := range ch { // it iterates until the channel is closed
		time.Sleep(time.Millisecond * 1000)
		println(v)
	}
	sig <- struct{}{}
}
