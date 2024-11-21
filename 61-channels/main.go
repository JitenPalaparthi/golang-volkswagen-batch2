package main

func main() {
	defer println("End of main")
	ch := make(chan int)
	//signal := make(chan bool) // bool -> 1 byte
	signal := make(chan struct{}) // bool -> 1 byte
	//fmt.Println("size of", unsafe.Sizeof(ch))
	go sendSq(ch, 20)
	go receiveSq(ch, signal)

	// v := <-signal // until the value is received
	// fmt.Println(v)
	<-signal
}

func sendSq(ch chan int, n int) {
	ch <- n * n
}

func receiveSq(ch chan int, sig chan struct{}) {
	println(<-ch)
	sig <- struct{}{}
}
