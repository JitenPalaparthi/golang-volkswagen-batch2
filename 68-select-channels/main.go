package main

import "time"

func main() {
	ch := make(chan int)
	sig1, sig2 := Receiver(ch, time.Second*10)
	go Generate(ch, sig2)
	<-sig1
}

func Generate(ch chan int, sig chan struct{}) {
	go func() {
		i := 1
	out:
		for {
			select {
			case ch <- i * i:
				time.Sleep(time.Millisecond * 500)
			case <-sig:
				close(ch)
				break out
			}
			i++
		}
	}()
}

func Receiver(ch1 chan int, d time.Duration) (chan struct{}, chan struct{}) {
	sig1 := make(chan struct{})
	sig2 := make(chan struct{})
	chtime := myTimeAfter(d) //time.After(d) // 10 seconds
	go func() {
	out:
		for {
			select {
			case v1, ok1 := <-ch1:
				if ok1 {
					println(v1)
				} else {
					break out
				}
			case <-chtime:
				sig2 <- struct{}{}
				sig1 <- struct{}{}
				close(sig1)
				close(sig2)
			}
		}
	}()

	return sig1, sig2
}

func myTimeAfter(d time.Duration) chan struct{} {
	sig := make(chan struct{})
	go func() {
		time.Sleep(d)
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}

// Generator
// Future
// Pipelines
// Fanin, Fanout
