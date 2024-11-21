package main

import "time"

func main() {
	chping, chpong := make(chan string), make(chan string)
	go ping(chping, chpong)
	<-pong(chping, chpong)
}

func ping(ch chan<- string, chpong <-chan string) {
	for {
		time.Sleep(time.Millisecond * 100)
		ch <- "ping"
		println(<-chpong)
	}
}

func pong(chping <-chan string, chpong chan<- string) <-chan struct{} {
	sig := make(chan struct{})
	go func() {
		for {
			println(<-chping)
			chpong <- "pong"
		}
	}()
	sig <- struct{}{}
	return sig
}
