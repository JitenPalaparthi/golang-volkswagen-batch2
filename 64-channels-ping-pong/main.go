package main

import (
	"fmt"
	"time"
)

func main() {
	chping, chpong := make(chan string), make(chan string)
	go pingS(chping, chpong)
	<-pong(chping, chpong)
}

func ping(ch chan<- string, chpong <-chan string) {
	i := 1
	for {
		time.Sleep(time.Millisecond * 1000)
		ch <- "ping-->" + fmt.Sprint(i)
		println(<-chpong)
		i++
	}
}

func pingS(chping chan<- string, chpong <-chan string) {
	i := 1
	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case chping <- "ping" + fmt.Sprint(i):
			println("Just sending ping thats all..")
		case v := <-chpong:
			println(v)
			// default:
			// 	println("---------------")
		}
		i++
	}
}

func pong(chping <-chan string, chpong chan<- string) <-chan struct{} {
	sig := make(chan struct{})
	go func() {
		i := 1
		for {
			println(<-chping)
			chpong <- "pong-->" + fmt.Sprint(i)
			i++
		}
	}()
	sig <- struct{}{}
	return sig
}
