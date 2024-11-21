package main

import "time"

func main() {
	ch := genSquares(100)
	//signal := receiver(ch)
	<-receiver(ch)
	//<-signal
}

// generator pattern
func genSquares(num int) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func receiver(ch <-chan int) chan struct{} {
	signal := make(chan struct{})
	go func() {
		for v := range ch {
			time.Sleep(time.Millisecond * 50)
			println(v)
		}
		signal <- struct{}{}
		close(signal)
	}()
	return signal
}
