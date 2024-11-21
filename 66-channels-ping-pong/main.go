package main

import (
	"fmt"
	"time"
)

// wrong ping and pong ..Dont use single channel even it looks it can send and receive data
func main() {
	chping := make(chan string)
	sig := make(chan struct{})
	go pong(chping, sig)
	go ping(chping)
	<-sig
}

func ping(pingch chan string) {
	i := 1
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case pingch <- "ping-->" + fmt.Sprint(i):
			i++
		case v := <-pingch:
			println("receive", v)
		}
	}
}

func pong(pongch chan string, sig chan struct{}) {
	i := 1
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case pongch <- "pong-->send ->" + fmt.Sprint(i):
			i++
		case v := <-pongch:
			println(v)
		}
	}
	sig <- struct{}{}
}
